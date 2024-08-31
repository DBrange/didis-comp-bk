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
        _id: new ObjectId("66c2804dcd58ba6c51446e02"),
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c28051cd58ba6c51446e03"),
        competitor_id: new ObjectId("66c27fddcd58ba6c51446dc5"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c28054cd58ba6c51446e04"),
        competitor_id: new ObjectId("66c27fe2cd58ba6c51446dca"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c28059cd58ba6c51446e05"),
        competitor_id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2805ccd58ba6c51446e06"),
        competitor_id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2805fcd58ba6c51446e07"),
        competitor_id: new ObjectId("66c27feccd58ba6c51446dd9"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c28062cd58ba6c51446e08"),
        competitor_id: new ObjectId("66c27fefcd58ba6c51446dde"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c28065cd58ba6c51446e09"),
        competitor_id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c07e6ed976cbb9359a5"),
        competitor_id: new ObjectId("66c60bd4e6ed976cbb93596c"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c0ae6ed976cbb9359a6"),
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935971"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c0de6ed976cbb9359a7"),
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935976"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c11e6ed976cbb9359a8"),
        competitor_id: new ObjectId("66c60bd6e6ed976cbb93597b"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c14e6ed976cbb9359a9"),
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935980"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c17e6ed976cbb9359aa"),
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935985"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c1be6ed976cbb9359ab"),
        competitor_id: new ObjectId("66c60bd8e6ed976cbb93598a"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60c1ee6ed976cbb9359ac"),
        competitor_id: new ObjectId("66c60bd9e6ed976cbb93598f"),
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
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
