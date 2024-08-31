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
    const collection = db.collection("guest_competitors");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66bfda60b2afd1e71e652ff7"),
        guest_user_id: new ObjectId("66bfda60b2afd1e71e652ff3"),
        competitor_id: new ObjectId("66bfda60b2afd1e71e652ff5"),
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
