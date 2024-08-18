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
        _id: new ObjectId("66b82564625966db1dc69b02"),
        user_id: new ObjectId("66b8252ee85d47dc5e737abc"),
        competitor_id: new ObjectId("66b82564625966db1dc69aff"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82567625966db1dc69b07"),
        user_id: new ObjectId("66b8253b625966db1dc69ae7"),
        competitor_id: new ObjectId("66b82567625966db1dc69b04"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8256c625966db1dc69b0c"),
        user_id: new ObjectId("66b8253d625966db1dc69aea"),
        competitor_id: new ObjectId("66b8256c625966db1dc69b09"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82572625966db1dc69b11"),
        user_id: new ObjectId("66b82540625966db1dc69aed"),
        competitor_id: new ObjectId("66b82572625966db1dc69b0e"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82576625966db1dc69b16"),
        user_id: new ObjectId("66b82543625966db1dc69af0"),
        competitor_id: new ObjectId("66b82576625966db1dc69b13"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8257a625966db1dc69b1b"),
        user_id: new ObjectId("66b82548625966db1dc69af3"),
        competitor_id: new ObjectId("66b8257a625966db1dc69b18"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8257e625966db1dc69b20"),
        user_id: new ObjectId("66b8254a625966db1dc69af6"),
        competitor_id: new ObjectId("66b8257e625966db1dc69b1d"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82581625966db1dc69b25"),
        user_id: new ObjectId("66b8254c625966db1dc69af9"),
        competitor_id: new ObjectId("66b82581625966db1dc69b22"),
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
