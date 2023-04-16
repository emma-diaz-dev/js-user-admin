export default () => ({
    port: parseInt(process.env.PORT, 10) || 5050,
    env: process.env.ENV || "local",
    database: {
      host: process.env.MONGO_DB_HOST || "localhost",
      port: parseInt(process.env.MONGO_DB_PORT, 10) || 27017
    }
  });