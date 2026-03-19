import { useState } from "react"

export default function Register() {

  const [name,setName] = useState("")
  const [phone,setPhone] = useState("")
  const [email,setEmail] = useState("")
  const [password,setPassword] = useState("")
  const [location,setLocation] = useState("")
  const [birthdate,setBirthdate] = useState("")
  const [role,setRole] = useState("user")

  // uber fields
  const [rid,setRID] = useState("")
  const [carPlate,setCarPlate] = useState("")
  const [carModel,setCarModel] = useState("")
  const [carColor,setCarColor] = useState("")
  const [carPhoto,setCarPhoto] = useState<File | null>(null)

  const [ridPhoto,setRidPhoto] = useState<File | null>(null)

  async function handleRegister() {

    let carId = null

    // =========================
    // 1 Register Car if uber
    // =========================
    if(role === "uber"){

      const carForm = new FormData()

      carForm.append("plate",carPlate)
      carForm.append("model",carModel)
      carForm.append("color",carColor)

      if(carPhoto){
        carForm.append("photo",carPhoto)
      }

      const carRes = await fetch("http://localhost:2534/api/v001/car/register",{
        method:"POST",
        body:carForm
      })

      const carData = await carRes.json()

      console.log("car response",carData)

      carId = carData.id
    }

    // =========================
    // 2 Register User
    // =========================

    const form = new FormData()

    form.append("name",name)
    form.append("phone",phone)
    form.append("role",role)
    form.append("email",email)
    form.append("password",password)
    form.append("location",location)
    form.append("birthdate",birthdate)

    if(role === "uber"){

      form.append("rid",rid)

      if(ridPhoto){
        form.append("photo",ridPhoto)
      }

      form.append("carId",carId)
    }

    const res = await fetch("http://localhost:2534/api/v001/register",{
      method:"POST",
      body:form
    })

    const data = await res.json()

    console.log("user response",data)

  }

  return (
    <div>

      <h1>Register</h1>

      <input placeholder="name" onChange={(e)=>setName(e.target.value)}/>
      <input placeholder="phone" onChange={(e)=>setPhone(e.target.value)}/>
      <input placeholder="email" onChange={(e)=>setEmail(e.target.value)}/>
      <input type="password" placeholder="password" onChange={(e)=>setPassword(e.target.value)}/>
      <input placeholder="location" onChange={(e)=>setLocation(e.target.value)}/>
      <input type="date" onChange={(e)=>setBirthdate(e.target.value)}/>

      <select onChange={(e)=>setRole(e.target.value)}>
        <option value="user">User</option>
        <option value="uber">Uber</option>
      </select>

      {role === "uber" && (

        <div>

          <h2>Uber Info</h2>

          <input placeholder="RID" onChange={(e)=>setRID(e.target.value)}/>
          <input type="file" onChange={(e)=>setRidPhoto(e.target.files?.[0] || null)}/>

          <h2>Car</h2>

          <input placeholder="plate" onChange={(e)=>setCarPlate(e.target.value)}/>
          <input placeholder="model" onChange={(e)=>setCarModel(e.target.value)}/>
          <input placeholder="color" onChange={(e)=>setCarColor(e.target.value)}/>
          <input type="file" onChange={(e)=>setCarPhoto(e.target.files?.[0] || null)}/>

        </div>

      )}

      <button onClick={handleRegister}>
        Register
      </button>

    </div>
  )
}