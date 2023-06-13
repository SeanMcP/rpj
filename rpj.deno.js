import { parse } from "https://deno.land/std@0.190.0/flags/mod.ts";

const args = parse(Deno.args);

function handleError(e, message) {
  if (args.verbose) {
    console.error(e);
  }
  console.log(message);
  Deno.exit(1);
}

const fieldMap = new Map([
  ["d", "description"],
  ["dep", "dependencies"],
  ["dev", "devDependencies"],
  ["n", "name"],
  ["peer", "peerDependencies"],
  ["v", "version"],
]);

const text = await Deno.readTextFile("package.json").catch((e) =>
  handleError(e, "No package.json file found in this directory")
);

try {
  const json = JSON.parse(text);
  const [field = "scripts"] = args._;
  const property = fieldMap.get(field) || field;
  console.log(json[property]);
} catch (e) {
  handleError(e, "Invalid package.json file");
}
