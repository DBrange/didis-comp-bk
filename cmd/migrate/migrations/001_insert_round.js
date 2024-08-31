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
        _id: new ObjectId("66c2802ccd58ba6c51446dea"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        round: "F",
        total_prize: 10000,
        points: 1000,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446deb"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        round: "SF",
        total_prize: 2000,
        points: 800,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dec"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        round: "CF",
        total_prize: 0,
        points: 0,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },

      // DOUBLE ELINATION
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df9"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        round: "F",
        total_prize: 400,
        points: 400,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dfa"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        round: "SF",
        total_prize: 100,
        points: 100,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb935996"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        round: "F",
        total_prize: 100,
        points: 100,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb935997"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        round: "SF",
        total_prize: 100,
        points: 100,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb935998"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        round: "CF",
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
