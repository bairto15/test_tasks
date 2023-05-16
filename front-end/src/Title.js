import React from "react";

export default function Title({ login, password, setRouter }) {
    async function handleOut() {
        try {
            const res = await fetch("http://localhost:8080/api", {
                method: "PUT",
                headers: { 
                    "Content-Type": "application/json;charset=utf-8",
                    "Authorization": login
                },
                body: JSON.stringify({ login, password })
            })

            if (res.status === 200) {
                setRouter("auth")
            }
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <div>
            <div className="quit" onClick={handleOut}>Выйти</div>
        </div>
    )
}