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
        _id: new ObjectId("66bcc1e583cd1ccac51d1122"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b82564625966db1dc69aff"),
        matches: [
          new ObjectId("66b8260e625966db1dc69b2b"),
          new ObjectId("66b8260e625966db1dc69b2b"),
          new ObjectId("66bd32e2b41cff099c32eeff"),
          new ObjectId("66bd3359b41cff099c32ef05"),
        ],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bcc1e583cd1ccac51d1123"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b82567625966db1dc69b04"),
        matches: [new ObjectId("66b8260e625966db1dc69b2b")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bcc1e583cd1ccac51d1124"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b8256c625966db1dc69b09"),
        matches: [
          new ObjectId("66b8260e625966db1dc69b31"),
          new ObjectId("66bd32e2b41cff099c32ef02"),
          new ObjectId("66bd3359b41cff099c32ef05"),
        ],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bcc1e583cd1ccac51d1125"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b82572625966db1dc69b0e"),
        matches: [
          new ObjectId("66b8260e625966db1dc69b2e"),
          new ObjectId("66bd32e2b41cff099c32eeff"),
        ],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bcc1e583cd1ccac51d1126"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b82576625966db1dc69b13"),
        matches: [new ObjectId("66b8260e625966db1dc69b31")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bcc1e583cd1ccac51d1127"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b8257a625966db1dc69b18"),
        matches: [new ObjectId("66b8260e625966db1dc69b2e")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bcc1e583cd1ccac51d1128"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b8257e625966db1dc69b1d"),
        matches: [
          new ObjectId("66b8260e625966db1dc69b34"),
          new ObjectId("66bd32e2b41cff099c32ef02"),
        ],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bcc1e583cd1ccac51d1129"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66b82581625966db1dc69b22"),
        matches: [new ObjectId("66b8260e625966db1dc69b34")],
        tournaments_won: [],
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda60b2afd1e71e652ff6"),
        total_wins: 0,
        total_losses: 0,
        money_earned: 0,
        competitor_id: new ObjectId("66bfda60b2afd1e71e652ff5"),
        matches: [new ObjectId("66b8260e625966db1dc69b34")],
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
