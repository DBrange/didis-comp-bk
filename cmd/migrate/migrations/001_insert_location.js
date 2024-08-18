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
        _id: new ObjectId("66bae5f38f7e1d13c7687467"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c7687468"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c7687469"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c768746a"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c768746b"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c768746c"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c768746d"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c768746e"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bae5f38f7e1d13c768746f"),
        state: "algun lugar lindo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfd9930b2ca89158341cd8"),
        state: "algun lugar del torneo",
        country: "algun pais lindo",
        city: "alguna ciudad linda",
        lat: null,
        long: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfd9dab2afd1e71e652fe4"),
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
