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
    const collection = db.collection("tournament_registrations");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66b8265f625966db1dc69b37"),
        competitor_id: new ObjectId("66b82564625966db1dc69aff"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82666625966db1dc69b38"),
        competitor_id: new ObjectId("66b82567625966db1dc69b04"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8266a625966db1dc69b39"),
        competitor_id: new ObjectId("66b8256c625966db1dc69b09"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8266f625966db1dc69b3a"),
        competitor_id: new ObjectId("66b82572625966db1dc69b0e"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82673625966db1dc69b3b"),
        competitor_id: new ObjectId("66b82576625966db1dc69b13"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82676625966db1dc69b3c"),
        competitor_id: new ObjectId("66b8257a625966db1dc69b18"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82679625966db1dc69b3d"),
        competitor_id: new ObjectId("66b8257e625966db1dc69b1d"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8267d625966db1dc69b3e"),
        competitor_id: new ObjectId("66b82581625966db1dc69b22"),
        tournament_id: new ObjectId("66b8260e625966db1dc69b28"),
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
