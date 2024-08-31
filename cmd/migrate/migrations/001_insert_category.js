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
    const collection = db.collection("categories");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c28019cd58ba6c51446de7"),
        name: "Categoria A",
        genre: "M",
        total_participants: 4,
        range_movement: "ONE_WEEK",
        average_score: 0,
        sport: "TENNIS",
        competitor_type: "S",
        organizer_id: new ObjectId("66c27fb7cd58ba6c51446dbe"),
        tournaments: [new ObjectId("66c2802ccd58ba6c51446de9")],
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
