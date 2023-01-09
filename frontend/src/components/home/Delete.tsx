import axios from 'axios'
import React from 'react'
import { HOST } from '../../App'

function articleDelete(id:number) {
    axios
		.post(`${HOST}/delete/${id}`)
		.then(res => {
			JSON.stringify(id)
		})
		.catch(() => {
			console.log("削除出来ませんでした")
		})
}

function Delete(props:any) {
  return (
    <button onClick={() => articleDelete(props.index)}>削除</button>
  )
}

export default Delete