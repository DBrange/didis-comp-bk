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
        _id: new ObjectId("66c2802ccd58ba6c51446de9"),
        name: "primer torneo que organizo",
        start_date: null,
        finish_date: null,
        points: 1200,
        total_prize: 12000,
        total_competitors: 8,
        max_capacity: 8,
        average_score: null,
        genre: "M",
        sport: "TENNIS",
        surface: "CLAY",
        competitor_type: "S",
        location_id: new ObjectId("66c2802ccd58ba6c51446de8"),
        organizer_id: new ObjectId("66c27fb7cd58ba6c51446dbe"),
        category_id: new ObjectId("66c28019cd58ba6c51446de7"),
        double_elimination_id: new ObjectId("66c2802ccd58ba6c51446e01"),
        rounds: [
          new ObjectId("66c2802ccd58ba6c51446dea"),
          new ObjectId("66c2802ccd58ba6c51446deb"),
          new ObjectId("66c2802ccd58ba6c51446dec"),
        ],
        matches: [
          new ObjectId("66c2802ccd58ba6c51446ded"),
          new ObjectId("66c2802ccd58ba6c51446df0"),
          new ObjectId("66c2802ccd58ba6c51446df3"),
          new ObjectId("66c2802ccd58ba6c51446df6"),
        ],
        pots: [],
        groups: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb935994"),
        name: "primer torneo que organizo",
        start_date: null,
        finish_date: null,
        points: 1200,
        total_prize: 12000,
        total_competitors: 8,
        max_capacity: 8,
        average_score: null,
        availability: {
          available_courts: 2,
          average_hours: 2,
        },
        genre: "M",
        sport: "TENNIS",
        surface: "CLAY",
        competitor_type: "S",
        location_id: new ObjectId("66c60bebe6ed976cbb935994"),
        organizer_id: new ObjectId("66c60bb40ac870e55afaa26d"),
        category_id: null,
        double_elimination_id: null,
        rounds: [
          new ObjectId("66c60bebe6ed976cbb935996"),
          new ObjectId("66c60bebe6ed976cbb935997"),
          new ObjectId("66c60bebe6ed976cbb935998"),
        ],
        matches: [
          new ObjectId("66c60bebe6ed976cbb935999"),
          new ObjectId("66c60bebe6ed976cbb93599c"),
          new ObjectId("66c60bebe6ed976cbb93599f"),
          new ObjectId("66c60bebe6ed976cbb9359a2"),
        ],
        pots: [],
        groups: [],
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
