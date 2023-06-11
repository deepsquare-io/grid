module.exports = {
  root: true,
  plugins: ['prettier'],
  extends: ['prettier', 'plugin:prettier/recommended', 'plugin:import/recommended'],
  rules: {
    'prettier/prettier': 'error',
    'no-unused-vars': ['error', { vars: 'all', args: 'after-used', ignoreRestSiblings: false }],
  },
  parserOptions: {
    ecmaVersion: 2018,
    sourceType: 'module',
  },
};
