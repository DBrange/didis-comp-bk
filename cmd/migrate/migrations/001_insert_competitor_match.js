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
    const collection = db.collection("competitor_matches");

    // Insertar datos de prueba
    await collection.insertMany([
      // TOURNAMENT BRACKET
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dee"),
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        match_id: new ObjectId("66c2802ccd58ba6c51446ded"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446def"),
        competitor_id: new ObjectId("66c27fddcd58ba6c51446dc5"),
        match_id: new ObjectId("66c2802ccd58ba6c51446ded"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df1"),
        competitor_id: new ObjectId("66c27fe2cd58ba6c51446dca"),
        match_id: new ObjectId("66c2802ccd58ba6c51446df0"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df2"),
        competitor_id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        match_id: new ObjectId("66c2802ccd58ba6c51446df0"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df4"),
        competitor_id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        match_id: new ObjectId("66c2802ccd58ba6c51446df3"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df5"),
        competitor_id: new ObjectId("66c27feccd58ba6c51446dd9"),
        match_id: new ObjectId("66c2802ccd58ba6c51446df3"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df7"),
        competitor_id: new ObjectId("66c27fefcd58ba6c51446dde"),
        match_id: new ObjectId("66c2802ccd58ba6c51446df6"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446df8"),
        competitor_id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        match_id: new ObjectId("66c2802ccd58ba6c51446df6"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dfc"),
        competitor_id: new ObjectId("66c27fddcd58ba6c51446dc5"),
        match_id: new ObjectId("66c2802ccd58ba6c51446dfb"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dfd"),
        competitor_id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        match_id: new ObjectId("66c2802ccd58ba6c51446dfb"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446dff"),
        competitor_id: new ObjectId("66c27feccd58ba6c51446dd9"),
        match_id: new ObjectId("66c2802ccd58ba6c51446dfe"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c2802ccd58ba6c51446e00"),
        competitor_id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        match_id: new ObjectId("66c2802ccd58ba6c51446dfe"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c358a791147609dcc5ef3a"),
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        match_id: new ObjectId("66c358a791147609dcc5ef39"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c358a791147609dcc5ef3b"),
        competitor_id: new ObjectId("66c27fe2cd58ba6c51446dca"),
        match_id: new ObjectId("66c358a791147609dcc5ef39"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c358a791147609dcc5ef3d"),
        competitor_id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        match_id: new ObjectId("66c358a791147609dcc5ef3c"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c358a791147609dcc5ef3e"),
        competitor_id: new ObjectId("66c27fefcd58ba6c51446dde"),
        match_id: new ObjectId("66c358a791147609dcc5ef3c"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c3596c91147609dcc5ef40"),
        competitor_id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        match_id: new ObjectId("66c3596c91147609dcc5ef3f"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c3596c91147609dcc5ef41"),
        competitor_id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        match_id: new ObjectId("66c3596c91147609dcc5ef3f"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c35b193586c2fb83cb16fc"),
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        match_id: new ObjectId("66c35b193586c2fb83cb16fb"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c35b193586c2fb83cb16fd"),
        competitor_id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        match_id: new ObjectId("66c35b193586c2fb83cb16fb"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb93599a"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb935999"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb93599b"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb935999"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb93599d"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb93599c"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb93599e"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb93599c"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb9359a0"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb93599f"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb9359a1"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb93599f"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb9359a3"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb9359a2"),
        position: 1,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb9359a4"),
        competitor_id: null,
        match_id: new ObjectId("66c60bebe6ed976cbb9359a2"),
        position: 2,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },

      // TOURNAMENT GROUP
    ]);

    console.log("Datos insertados");
  } finally {
    await client.close();
  }
}

run().catch(console.dir);
