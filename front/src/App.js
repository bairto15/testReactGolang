import React from "react"
import './App.css';

function App() {
  const [state, setState] = React.useState([{ id: 0, name: "", struct: [""] }])
  const url = "http://localhost:8000/table"

  function changeName(e, i) {
    const copy = JSON.parse(JSON.stringify(state))
    copy[i].name = e.target.value
    setState(copy)
  }

  function changeStruct(e, i, iStruct) {
    const copy = JSON.parse(JSON.stringify(state))
    copy[i].struct[iStruct] = e.target.value
    setState(copy)
  }
  
  function handleAdd(i) {
    const copy = JSON.parse(JSON.stringify(state))
    if (!copy[i].struct) {
      copy[i].struct = []
    }
    copy[i].struct.push("")
    setState(copy)
  }

  async function handleSave() {
    const res = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json;charset=utf-8" },
      body: JSON.stringify(state),
    });
    const response = await res.json();
    if (response.error) {
      alert(response.error)
    }
  }

  React.useEffect(() => {
    fetch(url).then(res => res.json()).then(res => res && setState(res))
  }, [])

  return (
    <div className="App">
      {state.map((item, i) => (
        <div key={i} className="container">
          <div className="block_table">
            <input
              placeholder="Нзвание таблицы"
              className="input_name_table"
              value={item.name}
              onChange={(e) => changeName(e, i)}
            />
            <button className="btn_save" onClick={() => handleAdd(i)}>
              Добавить столбец
            </button>
          </div>
          <div className="App">
            {item.struct && item.struct.map((item, iStruct) => (
              <input
                key={iStruct}
                placeholder="Нзвание столбца"
                className="input_name_row"
                value={item}
                onChange={(e) => changeStruct(e, i, iStruct)}
              />
            ))}
          </div>
        </div>
      ))}
      <button className="btn_save" onClick={handleSave}>
        Сохранить
      </button>
    </div>
  );
}

export default App;
