const fastify = require("fastify")();

fastify.get("/plaintext", (req, reply) => {
  reply
    .header("Content-Type", "text/plain")
    .code(200)
    .send("Hello, World!");
});

fastify.listen({ port: 8081, host: "0.0.0.0" }, (err, address) => {
  if (err) {
    throw err;
  }

  console.log(
    `Worker started and listening on ${address} ${new Date().toISOString()}`
  );
});
