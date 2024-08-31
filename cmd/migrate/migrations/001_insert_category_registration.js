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
    const collection = db.collection("category_registrations");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c28428995393b87a45f918"),
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        category_id: new ObjectId("66c28019cd58ba6c51446de7"),
        points: 0,
        registered_positions: [],
        current_position: null,
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
