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
    const collection = db.collection("organizers");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c27fb7cd58ba6c51446dbe"),
        categories: [new ObjectId("66c28019cd58ba6c51446de7")],
        average_score: 0,
        average_tournament_score: 0,
        total_categories: 0,
        total_tournaments: 0,
        user_id: new ObjectId("66c27fb7cd58ba6c51446dbd"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bb40ac870e55afaa26d"),
        categories: [],
        average_score: 0,
        average_tournament_score: 0,
        total_categories: 0,
        total_tournaments: 0,
        user_id: new ObjectId("66c60bb40ac870e55afaa26c"),
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
