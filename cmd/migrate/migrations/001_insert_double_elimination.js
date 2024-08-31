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
    const collection = db.collection("double_eliminations");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c2802ccd58ba6c51446e01"),
        matches: [
          new ObjectId("66c2802ccd58ba6c51446dfb"),
          new ObjectId("66c2802ccd58ba6c51446dfe"),
          new ObjectId("66c2802ccd58ba6c51446dfb"),
          new ObjectId("66c2802ccd58ba6c51446dfe"),
          new ObjectId("66c3596c91147609dcc5ef3f"),
        ],
        rounds: [
          new ObjectId("66c2802ccd58ba6c51446df9"),
          new ObjectId("66c2802ccd58ba6c51446dfa"),
        ],
        total_prize: 333,
        points: 333,
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
