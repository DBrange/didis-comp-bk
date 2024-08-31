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
    const collection = db.collection("locations");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c27f6f54cbea309e42280e"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f82cd58ba6c51446da4"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f87cd58ba6c51446da7"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f8acd58ba6c51446daa"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f94cd58ba6c51446dad"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f99cd58ba6c51446db0"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f9ccd58ba6c51446db3"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fa0cd58ba6c51446db6"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fa7cd58ba6c51446db9"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fb7cd58ba6c51446dbc"),
        state: "algun lugar del torneo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446de8"),
        state: "torneo de group",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bb40ac870e55afaa26b"),
        state: "torneo de group",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bbee6ed976cbb935968"),
        state: "torneo de group",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb935993"),
        state: "torneo de group",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
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
