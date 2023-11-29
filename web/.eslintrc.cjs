/**
 * @type {import("eslint").Linter.Config}
 */
const config = {
    ignorePatterns: [".eslintrc.*"],
    env: {
        browser: true,
        es2021: true,
    },
    extends: [
        "eslint:recommended",
        "plugin:@typescript-eslint/recommended",
        "plugin:react/recommended",
    ],
    parser: "@typescript-eslint/parser",
    parserOptions: {
        ecmaVersion: "latest",
        sourceType: "module",
    },
    plugins: ["@typescript-eslint", "react"],
    rules: {
        "react/react-in-jsx-scope": "off"
    },
    settings: {
        react: {
            version: "detect",
        },
    },
};

module.exports = config;
