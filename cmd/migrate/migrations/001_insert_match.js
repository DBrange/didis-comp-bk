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
    const collection = db.collection("matches");

    // Insertar datos de prueba
    await collection.insertMany([
      // TOURNAMENT BRACKET
      // CF
      {
        _id: new ObjectId("66c2802ccd58ba6c51446ded"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446dec"),
        result: "",
        winner: new ObjectId("66c27fdacd58ba6c51446dc0"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df0"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446dec"),
        result: "",
        winner: new ObjectId("66c27fe2cd58ba6c51446dca"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 2,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df3"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446dec"),
        result: "",
        winner: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 3,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df6"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446dec"),
        result: "",
        winner: new ObjectId("66c27fefcd58ba6c51446dde"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 4,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },

      // SF
      {
        _id: new ObjectId("66c358a791147609dcc5ef39"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446deb"),
        result: "",
        winner: new ObjectId("66c27fdacd58ba6c51446dc0"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c358a791147609dcc5ef3c"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446deb"),
        result: "",
        winner: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 2,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      // F
      {
        _id: new ObjectId("66c35b193586c2fb83cb16fb"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446dea"),
        result: "",
        winner: new ObjectId("66c27fdacd58ba6c51446dc0"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },

      // DOUBLE ELIMINATION
      // SF
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dfb"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446dfa"),
        result: "",
        winner: new ObjectId("66c27fddcd58ba6c51446dc5"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dfe"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446dfa"),
        result: "",
        winner: new ObjectId("66c27feccd58ba6c51446dd9"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 2,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      // F
      {
        _id: new ObjectId("66c3596c91147609dcc5ef3f"),
        sport: "TENNIS",
        round_id: new ObjectId("66c2802ccd58ba6c51446df9"),
        result: "",
        winner: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        tournament_id: new ObjectId("66c2802ccd58ba6c51446de9"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },

      // PRUEBA
      {
        _id: new ObjectId("66c60bebe6ed976cbb935999"),
        sport: "TENNIS",
        round_id: new ObjectId("66c60bebe6ed976cbb935998"),
        result: "",
        winner: null,
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb93599c"),
        sport: "TENNIS",
        round_id: new ObjectId("66c60bebe6ed976cbb935998"),
        result: "",
        winner: null,
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb93599f"),
        sport: "TENNIS",
        round_id: new ObjectId("66c60bebe6ed976cbb935998"),
        result: "",
        winner: null,
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        position: 1,
        date: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb9359a2"),
        sport: "TENNIS",
        round_id: new ObjectId("66c60bebe6ed976cbb935998"),
        result: "",
        winner: null,
        tournament_id: new ObjectId("66c60bebe6ed976cbb935994"),
        position: 1,
        date: null,
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
