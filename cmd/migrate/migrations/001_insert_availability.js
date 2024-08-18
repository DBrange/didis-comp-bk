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
    const collection = db.collection("availabilities");

    // Insertar datos de prueba
    await collection.insertMany([
      {
        _id: new ObjectId("66b8252ee85d47dc5e737abd"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b8252ee85d47dc5e737abc"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8253b625966db1dc69ae8"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b8253b625966db1dc69ae7"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8253d625966db1dc69aeb"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b8253d625966db1dc69aea"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82540625966db1dc69aee"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b82540625966db1dc69aed"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82543625966db1dc69af1"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b82543625966db1dc69af0"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82548625966db1dc69af4"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b82548625966db1dc69af3"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8254a625966db1dc69af7"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b8254a625966db1dc69af6"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8254c625966db1dc69afa"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: new ObjectId("66b8254c625966db1dc69af9"),
        competitor_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82564625966db1dc69b00"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b82564625966db1dc69aff"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82567625966db1dc69b05"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b82567625966db1dc69b04"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8256c625966db1dc69b0a"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b8256c625966db1dc69b09"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82572625966db1dc69b0f"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b82572625966db1dc69b0e"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82576625966db1dc69b14"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b82576625966db1dc69b13"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8257a625966db1dc69b19"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b8257a625966db1dc69b18"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b8257e625966db1dc69b1e"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b8257e625966db1dc69b1d"),
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66b82581625966db1dc69b23"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "MONDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "TUESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "WEDNESDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "THURSDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "FRIDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
          {
            day: "SATURDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "NOT_AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "NOT_AVAILABLE",
              },
            ],
          },
        ],
        user_id: null,
        competitor_id: new ObjectId("66b82581625966db1dc69b22"),
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