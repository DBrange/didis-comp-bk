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
    const collection = db.collection("competitor_stats");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66c27fdacd58ba6c51446dc2"),
        total_wins: 3,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        matches: [new ObjectId("66c2802ccd58ba6c51446ded")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fddcd58ba6c51446dc7"),
        total_wins: 3,
        total_losses: 1,
        money_earned: 0,
        competitor_id: new ObjectId("66c27fddcd58ba6c51446dc5"),
        matches: [new ObjectId("66c2802ccd58ba6c51446ded")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe2cd58ba6c51446dcc"),
        total_wins: 1,
        total_losses: 1,
        money_earned: 0,
        competitor_id: new ObjectId("66c27fe2cd58ba6c51446dca"),
        matches: [new ObjectId("66c2802ccd58ba6c51446df0")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe6cd58ba6c51446dd1"),
        total_wins: 1,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        matches: [new ObjectId("66c2802ccd58ba6c51446df0")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe9cd58ba6c51446dd6"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        matches: [new ObjectId("66c2802ccd58ba6c51446df3")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27feccd58ba6c51446ddb"),
        total_wins: 2,
        total_losses: 1,
        money_earned: 0,
        competitor_id: new ObjectId("66c27feccd58ba6c51446dd9"),
        matches: [new ObjectId("66c2802ccd58ba6c51446df3")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fefcd58ba6c51446de0"),
        total_wins: 1,
        total_losses: 1,
        money_earned: 0,
        competitor_id: new ObjectId("66c27fefcd58ba6c51446dde"),
        matches: [new ObjectId("66c2802ccd58ba6c51446df6")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27ff3cd58ba6c51446de5"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        matches: [new ObjectId("66c2802ccd58ba6c51446df6")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd4e6ed976cbb93596e"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd4e6ed976cbb93596c"),
        matches: [],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935973"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935971"),
        matches: [],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935978"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935976"),
        matches: [],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd6e6ed976cbb93597d"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd6e6ed976cbb93597b"),
        matches: [],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935982"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935980"),
        matches: [],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935987"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935985"),
        matches: [],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd8e6ed976cbb93598c"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd8e6ed976cbb93598c"),
        matches: [],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd9e6ed976cbb935991"),
        total_wins: 0,
        total_losses: 2,
        money_earned: 0,
        competitor_id: new ObjectId("66c60bd9e6ed976cbb93598f"),
        matches: [],
        tournaments_won: [],
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
