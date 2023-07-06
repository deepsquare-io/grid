module.exports = {
  root: true,
  plugins: ['prettier', 'deprecation', 'react'],
  extends: [
    'plugin:react/recommended',
    'plugin:react/jsx-runtime',
    'prettier',
    'plugin:prettier/recommended',
    'eslint:recommended',
  ],
  parser: '@babel/eslint-parser',
  rules: {
    'prettier/prettier': 'error',
    'no-console': ['error', { allow: ['error', 'warn', 'debug'] }],
    'no-restricted-imports': 'error',
    'deprecation/deprecation': 'error',
  },
  parserOptions: {
    requireConfigFile: false,
    ecmaFeatures: {
      jsx: true,
    },
    babelOptions: {
      babelrc: false,
      configFile: false,
    },
  },
};
