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
        _id: new ObjectId("66b8272e625966db1dc69b3f"),
        competitor_id: new ObjectId("66b82564625966db1dc69aff"),
        category_id: new ObjectId("66b825ef625966db1dc69b26"),
        points: 1000,
        registered_positions: [],
        current_position: 4,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82732625966db1dc69b40"),
        competitor_id: new ObjectId("66b82567625966db1dc69b04"),
        category_id: new ObjectId("66b825ef625966db1dc69b26"),
        points: 2000,
        registered_positions: [],
        current_position: 3,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82736625966db1dc69b41"),
        competitor_id: new ObjectId("66b8256c625966db1dc69b09"),
        category_id: new ObjectId("66b825ef625966db1dc69b26"),
        points: 3000,
        registered_positions: [],
        current_position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8273b625966db1dc69b42"),
        competitor_id: new ObjectId("66b82572625966db1dc69b0e"),
        category_id: new ObjectId("66b825ef625966db1dc69b26"),
        points: 4000,
        registered_positions: [],
        current_position: 1,
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
