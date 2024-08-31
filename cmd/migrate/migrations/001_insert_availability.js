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
        _id: new ObjectId("66c27f6f54cbea309e422810"),
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
        user_id: new ObjectId("66c27f6f54cbea309e42280f"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f82cd58ba6c51446da6"),
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
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f87cd58ba6c51446da9"),
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
        user_id: new ObjectId("66c27f87cd58ba6c51446da8"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f8acd58ba6c51446dac"),
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
        user_id: new ObjectId("66c27f8acd58ba6c51446dab"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f94cd58ba6c51446daf"),
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
        user_id: new ObjectId("66c27f94cd58ba6c51446dae"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f99cd58ba6c51446db2"),
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
        user_id: new ObjectId("66c27f99cd58ba6c51446db1"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27f9ccd58ba6c51446db5"),
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
        user_id: new ObjectId("66c27f9ccd58ba6c51446db4"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fa0cd58ba6c51446db8"),
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
        user_id: new ObjectId("66c27fa0cd58ba6c51446db7"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fa7cd58ba6c51446dbb"),
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
        competitor_id: new ObjectId("66c27fa7cd58ba6c51446dba"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fdacd58ba6c51446dc1"),
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
        competitor_id: new ObjectId("66c27fdacd58ba6c51446dc0"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fddcd58ba6c51446dc6"),
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
        competitor_id: new ObjectId("66c27fddcd58ba6c51446dc5"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe2cd58ba6c51446dcb"),
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
        competitor_id: new ObjectId("66c27fe2cd58ba6c51446dca"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe6cd58ba6c51446dd0"),
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
        competitor_id: new ObjectId("66c27fe6cd58ba6c51446dcf"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27fe9cd58ba6c51446dd5"),
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
        competitor_id: new ObjectId("66c27fe9cd58ba6c51446dd4"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27feccd58ba6c51446dda"),
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
        competitor_id: new ObjectId("66c27fefcd58ba6c51446ddf"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c27ff3cd58ba6c51446de4"),
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
        competitor_id: new ObjectId("66c27ff3cd58ba6c51446de3"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },

      // PRUEBA

      {
        _id: new ObjectId("66c60bbee6ed976cbb93596a"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        user_id: new ObjectId("66c60bbee6ed976cbb935969"),
        competitor_id: null,
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd4e6ed976cbb93596d"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd4e6ed976cbb93596c"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935972"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935971"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd5e6ed976cbb935977"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd5e6ed976cbb935976"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd6e6ed976cbb93597c"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd6e6ed976cbb93597b"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935981"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935980"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd7e6ed976cbb935986"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd7e6ed976cbb935985"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd8e6ed976cbb93598b"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd8e6ed976cbb93598a"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bd9e6ed976cbb935990"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: new ObjectId("66c60bd9e6ed976cbb93598f"),
        tournament_id: null,
        created_at: new Date("1990-01-01T00:00:00Z"),
        updated_at: new Date("1990-01-01T00:00:00Z"),
      },
      {
        _id: new ObjectId("66c60bebe6ed976cbb935995"),
        daily_availabilities: [
          {
            day: "SUNDAY",
            time_slots: [
              {
                time_slot: "00:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "01:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "02:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "03:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "04:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "05:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "06:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "07:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "08:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "09:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "10:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "11:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "12:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "13:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "14:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "15:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "16:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "17:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "18:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "19:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "20:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "21:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "22:00",
                status: "AVAILABLE",
              },
              {
                time_slot: "23:00",
                status: "AVAILABLE",
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
        competitor_id: null,
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
