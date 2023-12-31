import * as fs from "fs";
import { buildSchema, introspectionFromSchema } from "graphql";

const schemaFile = fs.readFileSync(
  "../schemas/sbatchapi/schema.graphqls",
  "utf8"
);

const introspection = introspectionFromSchema(
  buildSchema(schemaFile, {
    assumeValid: true,
  })
);

let jsonSchemaString = JSON.stringify(introspection, null, 2);

if (!fs.existsSync("./schemas")) {
  fs.mkdirSync("./schemas");
}
fs.writeFileSync("./schemas/job.schema.json", jsonSchemaString);
