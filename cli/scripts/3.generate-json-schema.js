import fs from 'fs';
import { buildSchema, introspectionFromSchema } from 'graphql';
import { fromIntrospectionQuery } from 'graphql-2-json-schema';

const schemaFile = fs.readFileSync('../schemas/sbatchapi/schema.graphqls', 'utf8');

const introspection = introspectionFromSchema(
  buildSchema(schemaFile, {
    assumeValid: true,
  }),
);

let schema = fromIntrospectionQuery(introspection);

// Only fetch the job properties.
schema.properties = schema.definitions.Job.properties;

const jsonSchemaString = JSON.stringify(schema, null, 2);

fs.writeFileSync('./job.schema.json', jsonSchemaString);
