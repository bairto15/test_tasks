import React from "react";
import Title from "./Title";

export default function App() {
  const [router, setRouter] = React.useState("auth")

  const [login, setLogin] = React.useState("")
  const [password, setPassword] = React.useState("")
  const [isUser, setIdUser] = React.useState("")

  const [variants, setVariants] = React.useState([])

  const [task, setTask] = React.useState({})
  const [answer, setAnswer] = React.useState("")
  const [idTest, setIdTest] = React.useState("")
  const [id_variant, setId_variant] = React.useState("")

  const [result, setResult] = React.useState("")

  React.useEffect(() => {
    if (router === "auth") {
      setIdUser("")
      setVariants([])
      setTask({})
      setAnswer("")
      setIdTest("")
      setId_variant("")
      setResult("")
    }
  }, [router])

  async function handleAuth() {
    console.log(login)
    try {
      const res = await fetch("http://localhost:8080/api", {
        method: "POST",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
          "Authorization": login
        },
        body: JSON.stringify({ login, password })
      })

      const response = await res.json();

      console.log(response)

      if (response?.variants) {
        setRouter("start")
        setVariants(response.variants)
      }
      if (response?.idUser) {
        setIdUser(response?.idUser)
      }
    } catch (e) {
      console.log(e)
    }
  }

  async function handleTask(idVariant) {
    idVariant && setId_variant(idVariant)
    try {
      const url = `http://localhost:8080/api?idTest=${idTest}&idUser=${isUser}&answer=${answer}&idTask=${task.Count ? task.Count + 1 : "1"}&idVariant=${idVariant ?? id_variant}&corrAnswer=${task.CorrectAnswer}`

      console.log(url)

      const res = await fetch(url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
          "Authorization": login
        }
      })
      setAnswer("")

      const response = await res.json();

      console.log(response)

      if (response?.task) {
        setRouter("test")
        setTask(response.task)
      } else if (response?.result) {
        setRouter("result")
        setResult(response.result)
      }

      if (response?.idTest) {
        setIdTest(response.idTest)
      }
    } catch (e) {
      console.log(e)
    }
  }

  switch (router) {
    case "auth":
      return (
        <div className='container_signin'>
          <div className="block_signin">
            <input
              placeholder="Логин"
              className="auth_input"
              value={login}
              onChange={(e) => setLogin(e.target.value)}
            />
            <input
              placeholder="Пароль"
              className="auth_input"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <button
              className="auth_btn"
              onClick={handleAuth}
            >
              Войти
            </button>
          </div>
        </div>
      )

    case "start":
      return (
        <div className="container_start">
          <Title login={login} password={password} setRouter={(e) => setRouter(e)} />
          {variants?.map((item) => (
            <div className="variant" key={item.Id} onClick={() => handleTask(item.Id)}>
              {item.Name}
            </div>
          ))}
        </div>
      )

    case "test":
      return (
        <div className="container_task">
          <Title login={login} password={password} setRouter={(e) => setRouter(e)} />
          <div>{task.Task}</div>
          <div>
            <input
              type="radio"
              name="radio"
              value=""
              onChange={() => setAnswer(task.Answer1)}
              checked={answer === task.Answer1}
            />
            {task.Answer1}
          </div>
          <div>
            <input
              type="radio"
              name="radio"
              value=""
              onChange={() => setAnswer(task.Answer2)}
              checked={answer === task.Answer2}
            />
            {task.Answer2}
          </div>
          <div>
            <input
              type="radio"
              name="radio"
              value=""
              onChange={() => setAnswer(task.Answer3)}
              checked={answer === task.Answer3}
            />
            {task.Answer3}
          </div>
          <div>
            <input
              type="radio"
              name="radio"
              value=""
              onChange={() => setAnswer(task.Answer4)}
              checked={answer === task.Answer4}
            />
            {task.Answer4}
          </div>
          <button onClick={() => handleTask()}>Выбрать</button>
        </div>
      )

    case "result":
      return (
        <div className="container_task">
          <Title login={login} password={password} setRouter={(e) => setRouter(e)} />
          Вы правильно ответили на {result.Percent}%
        </div>
      )

    default:
      return <div>Не найдена страница</div>
  }
}

