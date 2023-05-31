import fs from "fs";
import path from "path";

const __dirname = path.resolve();
const filepath = path.join(__dirname, "package.json");

const [field = "scripts", ...flags] = process.argv.slice(2);

const fieldMap = new Map([
  ["d", "description"],
  ["dep", "dependencies"],
  ["dev", "devDependencies"],
  ["n", "name"],
  ["peer", "peerDependencies"],
  ["v", "version"],
]);

const isVerbose = flags.includes("--verbose");

fs.readFile(filepath, "utf8", (err, data) => {
  if (err) {
    isVerbose && console.log(err);
    console.log("No package.json file found in this directory");
    process.exit(1);
  }

  try {
    const json = JSON.parse(data);
    const property = fieldMap.get(field) || field;
    console.log(json[property]);
  } catch (e) {
    isVerbose && console.log(e);
    console.log("Invalid package.json file");
    process.exit(1);
  }
});
