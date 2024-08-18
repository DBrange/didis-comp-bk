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
      {
        _id: new ObjectId("66bfda62b2afd1e71e652ffd"),
        guest_user_id: new ObjectId("66bfda62b2afd1e71e652ff9"),
        competitor_id: new ObjectId("66bfda62b2afd1e71e652ffb"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda65b2afd1e71e653003"),
        guest_user_id: new ObjectId("66bfda65b2afd1e71e652fff"),
        competitor_id: new ObjectId("66bfda65b2afd1e71e653001"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda67b2afd1e71e653009"),
        guest_user_id: new ObjectId("66bfda67b2afd1e71e653005"),
        competitor_id: new ObjectId("66bfda67b2afd1e71e653007"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6ab2afd1e71e65300f"),
        guest_user_id: new ObjectId("66bfda6ab2afd1e71e65300b"),
        competitor_id: new ObjectId("66bfda6ab2afd1e71e65300d"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6cb2afd1e71e653015"),
        guest_user_id: new ObjectId("66bfda6cb2afd1e71e653011"),
        competitor_id: new ObjectId("66bfda6cb2afd1e71e653013"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6fb2afd1e71e65301b"),
        guest_user_id: new ObjectId("66bfda6fb2afd1e71e653017"),
        competitor_id: new ObjectId("66bfda6fb2afd1e71e653019"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda72b2afd1e71e653021"),
        guest_user_id: new ObjectId("66bfda72b2afd1e71e65301d"),
        competitor_id: new ObjectId("66bfda72b2afd1e71e65301f"),
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
