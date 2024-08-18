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
        _id: new ObjectId("66b82564625966db1dc69aff"),
        single_id: new ObjectId("66b82564625966db1dc69afe"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82567625966db1dc69b04"),
        single_id: new ObjectId("66b82567625966db1dc69b03"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8256c625966db1dc69b09"),
        single_id: new ObjectId("66b8256c625966db1dc69b08"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82572625966db1dc69b0e"),
        single_id: new ObjectId("66b82572625966db1dc69b0d"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82576625966db1dc69b13"),
        single_id: new ObjectId("66b82576625966db1dc69b12"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8257a625966db1dc69b18"),
        single_id: new ObjectId("66b8257a625966db1dc69b17"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8257e625966db1dc69b1d"),
        single_id: new ObjectId("66b8257e625966db1dc69b1c"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82581625966db1dc69b22"),
        single_id: new ObjectId("66b82581625966db1dc69b21"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda60b2afd1e71e652ff5"),
        single_id: new ObjectId("66bfda60b2afd1e71e652ff4"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda62b2afd1e71e652ffb"),
        single_id: new ObjectId("66bfda62b2afd1e71e652ffa"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda65b2afd1e71e653001"),
        single_id: new ObjectId("66bfda65b2afd1e71e653000"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda67b2afd1e71e653007"),
        single_id: new ObjectId("66bfda67b2afd1e71e653006"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6ab2afd1e71e65300d"),
        single_id: new ObjectId("66c2170ad9b18faf967efbdf"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6cb2afd1e71e653013"),
        single_id: new ObjectId("66bfda6cb2afd1e71e653012"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda6fb2afd1e71e653019"),
        single_id: new ObjectId("66bfda6fb2afd1e71e653018"),
        double_id: null,
        team_id: null,
        sport: "TENNIS",
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66bfda72b2afd1e71e65301f"),
        single_id: new ObjectId("66bfda72b2afd1e71e65301e"),
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
