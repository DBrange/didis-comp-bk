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
      {
        _id: new ObjectId("66bfda62b2afd1e71e652ff9"),
        first_name: "El invita2",
        last_name: "invita",
        email: "invitado@gmail.com",
        image: "path/to/image.jpg",
        genre: "M",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda65b2afd1e71e652fff"),
        first_name: "El invita3",
        last_name: "invita",
        email: "invitado@gmail.com",
        image: "path/to/image.jpg",
        genre: "M",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda67b2afd1e71e653005"),
        first_name: "El invita4",
        last_name: "invita",
        email: "invitado@gmail.com",
        image: "path/to/image.jpg",
        genre: "M",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6ab2afd1e71e65300b"),
        first_name: "El invita5",
        last_name: "invita",
        email: "invitado@gmail.com",
        image: "path/to/image.jpg",
        genre: "M",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6cb2afd1e71e653011"),
        first_name: "El invita6",
        last_name: "invita",
        email: "invitado@gmail.com",
        image: "path/to/image.jpg",
        genre: "M",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6fb2afd1e71e653017"),
        first_name: "El invita7",
        last_name: "invita",
        email: "invitado@gmail.com",
        image: "path/to/image.jpg",
        genre: "M",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda72b2afd1e71e65301d"),
        first_name: "El invita8",
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
