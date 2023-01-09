import axios from 'axios'
import React, { useState } from 'react'
import { HOST } from '../../App'
import './Board.css'

function Create() {
  const [content, setContent] = useState({content: "", error: ""})
  const handleChange = (e:any) => {
    setContent({content: e.target.value, error: ""})
  }
  
  const test = (e:any) => {
    e.preventDefault()
    axios
    .post(`${HOST}/create/articles`, {
      content: content.content
    })
    .then(() => {
      setContent({content: "", error: "投稿しました"})
    })
    .catch(() => {
      setContent({content: content.content, error: "投稿できませんでした"})
    })
  } 

  
  return (
    <div>
        <div className="form-err-msg">
            {content.error}
        </div>
        <form onSubmit={test}>
            <input type="text" value={content.content} onChange={handleChange}/>
            <button type='submit'>投稿</button>
        </form>
    </div>
  )
}

export default Create