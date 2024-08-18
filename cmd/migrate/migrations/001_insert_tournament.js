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
    const collection = db.collection("tournaments");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66b8260e625966db1dc69b28"),
        name: "primer torneo que organizo",
        start_date: null,
        finish_date: null,
        points: 1000,
        total_prize: 12000,
        total_competitors: 8,
        max_capacity: 8,
        average_score: null,
        genre: "M",
        sport: "TENNIS",
        surface: "CLAY",
        competitor_type: "S",
        location_id: new ObjectId("66bb733a5aba6f564e298afe"),
        organizer_id: new ObjectId("66b82550625966db1dc69afd"),
        category_id: new ObjectId("66b825ef625966db1dc69b26"),
        double_elimination_id: null,
        rounds: [
          new ObjectId("66b8260e625966db1dc69b2a"),
          new ObjectId("66b972b5fecbd064c2fa890b"),
          new ObjectId("66b8260e625966db1dc69b29"),
        ],
        matches: [
          new ObjectId("66b8260e625966db1dc69b2b"),
          new ObjectId("66b8260e625966db1dc69b2e"),
          new ObjectId("66b8260e625966db1dc69b31"),
          new ObjectId("66b8260e625966db1dc69b34"),
          new ObjectId("66bd3359b41cff099c32ef05"),
          new ObjectId("66bd32e2b41cff099c32ef02"),
          new ObjectId("66bd32e2b41cff099c32eeff"),
        ],
        pots: [],
        groups: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfd9dab2afd1e71e652fe5"),
        name: "torneo de grupos",
        start_date: null,
        finish_date: null,
        points: 2000,
        total_prize: 20000,
        total_competitors: 16,
        max_capacity: 16,
        average_score: null,
        genre: "M",
        sport: "TENNIS",
        surface: "CLAY",
        competitor_type: "S",
        location_id: new ObjectId("66bfd9dab2afd1e71e652fe4"),
        organizer_id: new ObjectId("66b82550625966db1dc69afd"),
        category_id: new ObjectId("66b825ef625966db1dc69b26"),
        double_elimination_id: null,
        rounds: [new ObjectId("66bfd9dab2afd1e71e652fea")],
        matches: [
          new ObjectId("66bfdcd35314e8f3de21c45a"),
          new ObjectId("66bfdcd35314e8f3de21c45d"),
          new ObjectId("66bfdcd35314e8f3de21c460"),
          new ObjectId("66bfdcd35314e8f3de21c463"),
          new ObjectId("66bfdcd35314e8f3de21c466"),
          new ObjectId("66bfdcd35314e8f3de21c469"),
          new ObjectId("66bfdcd35314e8f3de21c46c"),
          new ObjectId("66bfdcd35314e8f3de21c46f"),
          new ObjectId("66bfdcd35314e8f3de21c472"),
          new ObjectId("66bfdcd35314e8f3de21c475"),
          new ObjectId("66bfdcd35314e8f3de21c478"),
          new ObjectId("66bfdcd35314e8f3de21c47b"),
          new ObjectId("66bfdcd35314e8f3de21c47e"),
          new ObjectId("66bfdcd35314e8f3de21c481"),
          new ObjectId("66bfdcd35314e8f3de21c484"),
          new ObjectId("66bfdcd35314e8f3de21c487"),
          new ObjectId("66bfdcd35314e8f3de21c48a"),
          new ObjectId("66bfdcd35314e8f3de21c48d"),
          new ObjectId("66bfdcd35314e8f3de21c490"),
          new ObjectId("66bfdcd35314e8f3de21c493"),
          new ObjectId("66bfdcd35314e8f3de21c496"),
          new ObjectId("66bfdcd35314e8f3de21c499"),
          new ObjectId("66bfdcd35314e8f3de21c49c"),
          new ObjectId("66bfdcd35314e8f3de21c49f"),
        ],
        pots: [],
        groups: [
          new ObjectId("66bfd9dab2afd1e71e652fe6"),
          new ObjectId("66bfd9dab2afd1e71e652fe7"),
          new ObjectId("66bfd9dab2afd1e71e652fe8"),
          new ObjectId("66bfd9dab2afd1e71e652fe9"),
        ],
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
