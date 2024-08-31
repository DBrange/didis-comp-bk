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
    const collection = db.collection("competitors");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        single_id: new ObjectId("66c27fdacd58ba6c51446dbf"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fddcd58ba6c51446dc5"),
        single_id: new ObjectId("66c27fddcd58ba6c51446dc4"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe2cd58ba6c51446dca"),
        single_id: new ObjectId("66c27fe2cd58ba6c51446dc9"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        single_id: new ObjectId("66c27fe6cd58ba6c51446dce"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        single_id: new ObjectId("66c27fe9cd58ba6c51446dd3"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27feccd58ba6c51446dd9"),
        single_id: new ObjectId("66c27feccd58ba6c51446dd8"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fefcd58ba6c51446dde"),
        single_id: new ObjectId("66c27fefcd58ba6c51446ddd"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        single_id: new ObjectId("66c27ff3cd58ba6c51446de2"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd4e6ed976cbb93596c"),
        single_id: new ObjectId("66c60bd4e6ed976cbb93596b"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935971"),
        single_id: new ObjectId("66c60bd5e6ed976cbb935970"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935976"),
        single_id: new ObjectId("66c60bd5e6ed976cbb935975"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd6e6ed976cbb93597b"),
        single_id: new ObjectId("66c60bd6e6ed976cbb93597a"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935980"),
        single_id: new ObjectId("66c60bd7e6ed976cbb93597f"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935985"),
        single_id: new ObjectId("66c60bd7e6ed976cbb935984"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd8e6ed976cbb93598a"),
        single_id: new ObjectId("66c60bd8e6ed976cbb935989"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd9e6ed976cbb93598f"),
        single_id: new ObjectId("66c60bd9e6ed976cbb93598e"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
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
