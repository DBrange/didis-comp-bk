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
    const collection = db.collection("guest_users");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66bfda60b2afd1e71e652ff3"),
        first_name: "El invita1",
        last_name: "invita",
        email: "invitado@gmail.com",
        image: "path/to/image.jpg",
        genre: "M",
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
