export default {
  clearMocks: true,
  coverageProvider: "v8",
  moduleFileExtensions: ["js", "ts", "json", "node"],

  roots: ["<rootDir>"],

  testMatch: ["**/*.test.[jt]s?(x)"],
  transform: {
    "^.+\\.(ts)$": "ts-jest",
  },
};
