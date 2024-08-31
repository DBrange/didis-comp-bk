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
    const collection = db.collection("tournament_groups");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66bfd9dab2afd1e71e652fe6"),
        tournament_id: new ObjectId("66bfd9dab2afd1e71e652fe5"),
        competitors: [
          {
            competitor_id: new ObjectId("66bfd9dab2afd1e71e652fe5"),
            stats: {
              matches_played: 0,
              matches_lost: 0,
              matches_won: 0,
              sets_won: 0,
              sets_lost: 0,
              games_won: 0,
              games_lost: 0,
            },
          },
          {
            competitor_id: new ObjectId("66b82567625966db1dc69b04"),
            stats: {
              matches_played: 0,
              matches_lost: 0,
              matches_won: 0,
              sets_won: 0,
              sets_lost: 0,
              games_won: 0,
              games_lost: 0,
            },
          },
          {
            competitor_id: new ObjectId("66b8256c625966db1dc69b09"),
            stats: {
              matches_played: 0,
              matches_lost: 0,
              matches_won: 0,
              sets_won: 0,
              sets_lost: 0,
              games_won: 0,
              games_lost: 0,
            },
          },
          {
            competitor_id: new ObjectId("66b82572625966db1dc69b0e"),
            stats: {
              matches_played: 0,
              matches_lost: 0,
              matches_won: 0,
              sets_won: 0,
              sets_lost: 0,
              games_won: 0,
              games_lost: 0,
            },
          },
        ],
        matches: [
          new ObjectId("66bfdcd35314e8f3de21c45a"),
          new ObjectId("66bfdcd35314e8f3de21c45d"),
          new ObjectId("66bfdcd35314e8f3de21c460"),
          new ObjectId("66bfdcd35314e8f3de21c463"),
          new ObjectId("66bfdcd35314e8f3de21c466"),
          new ObjectId("66bfdcd35314e8f3de21c469"),
        ],
        position: 0,
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
