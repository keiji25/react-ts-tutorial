import axios from 'axios'
import React, { useEffect, useReducer } from 'react'
import { HOST } from '../../App'
import './Board.css'
import Create from './Create'
import Delete from './Delete'

const inisialState = {
	loading :true,
	error: '',
	post: {}
}

const reducer = (state:any, action:any) => {
	switch(action.type) {
		case 'SUCCESS':
			return {
				loading: false,
				post: action.payload,
				error: ''
			}
		case 'ERROR':
			return {
				loading: false,
				post: {},
				error: "データ取得失敗"
			}
		default:
			return state
	}
}



function Board() {
	const [state, dispatch] = useReducer(reducer, inisialState)
	
	useEffect(() => {
		axios
		.get(`${HOST}/get/articles`)
		.then(res => {
			dispatch({type: 'SUCCESS', payload: res.data})
		})
		.catch(() => {
			dispatch({type: 'ERROR'})
		})
	}, [state])
	
  return (
    <div className='board'>
			<Create />
				{state.error 
					? <div className="err_message">{state.error}</div>
					: null
				}
                {state.loading ? "loading..." : null}
			<div className="board_content">
				{Object.keys(state.post).map((key:string, index:number) => {
						return (
							<div className="article_outline" key={index}>
								<p className="content">{state.post[key]['content']}</p>
                                <Delete index={state.post[key]['ID']} />
								<br />
							</div>
						)
					})
				}
			</div>		
		</div>
  )
}

export default Board