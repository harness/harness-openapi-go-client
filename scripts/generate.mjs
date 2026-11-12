#!/usr/bin/env node
// Regenerates a Go client with swagger-codegen (GoClientCodegen).
// Usage:
//   npm run generate
//   npm run generate -- nextgen
//   node scripts/generate.mjs nextgen --spec path/to/spec.yaml

import { spawnSync } from "node:child_process";
import { createInterface } from "node:readline/promises";
import { stdin as input, stdout as output } from "node:process";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const root = resolve(dirname(fileURLToPath(import.meta.url)), "..");

// swagger-codegen 3.0.33 matches nextgen/.swagger-codegen/VERSION.
const codegenVersion = "3.0.33";
const dockerImage = `swaggerapi/swagger-codegen-cli:${codegenVersion}`;

const clients = {
  nextgen: {
    spec: "nextgen/api/swagger.yaml",
    output: "nextgen",
    packageName: "nextgen",
  },
};

function usage() {
  console.log(`Usage: npm run generate -- <client> [--spec <file>]

Clients:
${Object.entries(clients)
  .map(([name, client]) => `  ${name}    spec ${client.spec}    output ${client.output}`)
  .join("\n")}
`);
}

function parseArgs(argv) {
  const positional = [];
  let spec;
  for (let i = 0; i < argv.length; i++) {
    const arg = argv[i];
    if (arg === "--help" || arg === "-h") {
      usage();
      process.exit(0);
    }
    if (arg === "--spec") {
      spec = argv[++i];
      if (!spec) {
        console.error("--spec needs a file path");
        process.exit(1);
      }
      continue;
    }
    positional.push(arg);
  }
  return { name: positional[0], spec };
}

async function pickClient() {
  const names = Object.keys(clients);
  console.log("Select the client to generate:\n");
  names.forEach((name, index) => {
    console.log(`  ${index + 1}) ${name}    ${clients[name].spec}`);
  });
  console.log("");
  const rl = createInterface({ input, output });
  const answer = (await rl.question("Client number: ")).trim();
  rl.close();
  const index = Number(answer) - 1;
  if (!names[index]) {
    console.error("No client selected.");
    process.exit(1);
  }
  return names[index];
}

function run(cmd, args) {
  const result = spawnSync(cmd, args, { stdio: "inherit" });
  if (result.error) {
    console.error(result.error.message);
    process.exit(1);
  }
  if (result.status !== 0) {
    process.exit(result.status ?? 1);
  }
}

function hasCommand(cmd) {
  const result = spawnSync("which", [cmd], { stdio: "ignore" });
  return result.status === 0;
}

function generate(name, specOverride) {
  const client = clients[name];
  if (!client) {
    console.error(`Unknown client "${name}".`);
    usage();
    process.exit(1);
  }

  const spec = specOverride ? resolve(specOverride) : join(root, client.spec);
  const out = join(root, client.output);
  const additional = [
    `packageName=${client.packageName}`,
    "packageVersion=1.0.0",
    "hideGenerationTimestamp=true",
  ].join(",");

  console.log(`Generating ${name}`);
  console.log(`  spec:   ${spec}`);
  console.log(`  output: ${out}`);
  console.log(`  generator: swagger-codegen ${codegenVersion} (go)`);

  if (hasCommand("docker")) {
    const specInContainer = specOverride ? "/local/input-spec.yaml" : `/local/${client.spec}`;
    const outInContainer = `/local/${client.output}`;
    const args = ["run", "--rm", "-v", `${root}:/local`];
    if (specOverride) {
      args.push("-v", `${spec}:/local/input-spec.yaml:ro`);
    }
    args.push(
      dockerImage,
      "generate",
      "-i",
      specInContainer,
      "-l",
      "go",
      "-o",
      outInContainer,
      "--additional-properties",
      additional,
    );
    run("docker", args);
    return;
  }

  if (hasCommand("swagger-codegen")) {
    console.log("Docker is not available. Using swagger-codegen on PATH.");
    run("swagger-codegen", [
      "generate",
      "-i",
      spec,
      "-l",
      "go",
      "-o",
      out,
      "--additional-properties",
      additional,
    ]);
    return;
  }

  console.error(
    `Install Docker, or install swagger-codegen ${codegenVersion} and put it on PATH.`,
  );
  process.exit(1);
}

const { name, spec } = parseArgs(process.argv.slice(2));
const selected = name || (await pickClient());
generate(selected, spec);
