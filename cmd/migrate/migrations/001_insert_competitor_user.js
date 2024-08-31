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
    const collection = db.collection("competitor_users");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c27fdacd58ba6c51446dc3"),
        user_id: new ObjectId("66c27f6f54cbea309e42280f"),
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fddcd58ba6c51446dc8"),
        user_id: new ObjectId("66c27f82cd58ba6c51446da5"),
        competitor_id: new ObjectId("66c27fddcd58ba6c51446dc5"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe2cd58ba6c51446dcd"),
        user_id: new ObjectId("66c27f87cd58ba6c51446da8"),
        competitor_id: new ObjectId("66c27fe2cd58ba6c51446dca"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe6cd58ba6c51446dd2"),
        user_id: new ObjectId("66c27f8acd58ba6c51446dab"),
        competitor_id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe9cd58ba6c51446dd7"),
        user_id: new ObjectId("66c27f94cd58ba6c51446dae"),
        competitor_id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27feccd58ba6c51446ddc"),
        user_id: new ObjectId("66c27f99cd58ba6c51446db1"),
        competitor_id: new ObjectId("66c27feccd58ba6c51446dd9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fefcd58ba6c51446de1"),
        user_id: new ObjectId("66c27f9ccd58ba6c51446db4"),
        competitor_id: new ObjectId("66c27fefcd58ba6c51446dde"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27ff3cd58ba6c51446de6"),
        user_id: new ObjectId("66c27fa0cd58ba6c51446db7"),
        competitor_id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd4e6ed976cbb93596f"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd4e6ed976cbb93596c"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935974"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935971"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935979"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935976"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd6e6ed976cbb93597e"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd6e6ed976cbb93597b"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935983"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935980"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935988"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935985"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd8e6ed976cbb93598d"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd8e6ed976cbb93598a"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd9e6ed976cbb935992"),
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: new ObjectId("66c60bd9e6ed976cbb93598f"),
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
