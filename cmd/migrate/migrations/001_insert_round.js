require("dotenv").config(); // Cargar las variables de entorno
const { MongoClient, ObjectId } = require("mongodb");

const uri = process.env.MONGO_URI; // Leer URI de MongoDB desde las variables de entorno
const client = new MongoClient(uri, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

async function run() {
  try {
    await client.connect();
    const db = client.db(process.env.DB_NAME);
    const collection = db.collection("rounds");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66b8260e625966db1dc69b29"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        round: "F",
        total_prize: 10000,
        points: 1000,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b972b5fecbd064c2fa890b"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        round: "SF",
        total_prize: 2000,
        points: 800,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8260e625966db1dc69b2a"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        round: "CF",
        total_prize: 1500,
        points: 500,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfd9dab2afd1e71e652fea"),
        tournament_id: new ObjectId("66bfd9dab2afd1e71e652fe5"),
        round: "GROUP",
        total_prize: 100,
        points: 100,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
    ]);

    console.log("Datos insertados");
  } finally {
    await client.close();
  }
}

run().catch(console.dir);
