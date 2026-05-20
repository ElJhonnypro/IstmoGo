import { useState } from "react"

export default function Login() {
    const [identifier,setIdentifier] = useState("")
    const [password,setPassword] = useState("")

    async function LoginHandler() {
        console.log("Login with",identifier)

        const res = await fetch("http://localhost:2534/api/v001/auth/login",{
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                identifier: identifier,
                password: password
            })
        })
        

        const data = await res.json()
        console.log("Login response",data)

        localStorage.setItem("token",data.Token)
        


    }

    return (
        <div>
            <h1>Login</h1>
            <input type="text" placeholder="ANY identifier" value={identifier} onChange={(e) => setIdentifier(e.target.value)} />
            <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
            <button onClick={LoginHandler}>Login</button>
        </div>
    )
}