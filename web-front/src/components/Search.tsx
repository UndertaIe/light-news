import React from 'react'
import { RouteComponentProps } from '@reach/router'
import 'antd/dist/antd.min.css'
import { Input } from 'antd'
import "../app.css"


interface Props extends RouteComponentProps {
  onSearch?: ((value: string, event?: React.KeyboardEvent<HTMLInputElement> | React.ChangeEvent<HTMLInputElement> | React.MouseEvent<HTMLElement, MouseEvent> | undefined) => void) | undefined
  onChange?:  React.ChangeEventHandler<HTMLInputElement> | undefined
  loading?: boolean
  
}
const { Search } = Input;
const SearchInput: React.FC<Props> = (props) => {
  return (
      <Search 
      className="search" 
      placeholder="输入关键字" 
      loading={props.loading} 
      onSearch={props.onSearch}
      onChange={props.onChange}
      allowClear={true}
      enterButton />
  )
}

export default SearchInput