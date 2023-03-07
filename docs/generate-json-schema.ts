import { introspectionFromSchema } from 'graphql';
import { GraphQLFileLoader } from '@graphql-tools/graphql-file-loader';
import { loadSchemaSync } from '@graphql-tools/load';
import fs from 'fs';

const options = {
  ignoreInternals: true,
  nullableArrayItems: true,
};

const schema = loadSchemaSync('../schemas/sbatchapi/schema.graphqls', {
  loaders: [new GraphQLFileLoader()],
});
const introspection = introspectionFromSchema(schema);

let jsonSchemaString = JSON.stringify(introspection, null, 2);

if (!fs.existsSync('./schemas')) {
  fs.mkdirSync('./schemas');
}
fs.writeFileSync('./schemas/job.schema.json', jsonSchemaString);
