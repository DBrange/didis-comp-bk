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
        _id: new ObjectId("66b825ef625966db1dc69b26"),
        name: "Categoria A",
        genre: "M",
        total_participants: 4,
        range_movement: "ONE_WEEK",
        average_score: 0,
        sport: "TENNIS",
        competitor_type: "S",
        organizer_id: new ObjectId("66b82550625966db1dc69afd"),
        tournaments: [new ObjectId("66b8260e625966db1dc69b28")],
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
