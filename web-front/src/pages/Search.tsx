/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, useEffect } from "react"
import { RouteComponentProps } from "@reach/router"

import { message } from "antd"
import "antd/dist/antd.min.css"

import SearchInput from "../components/Search"
import SearchService from "../services/search"
import NewsList, {NewsItem} from "../components/NewsList"

const Search: React.FC<RouteComponentProps> = (props) => {
  let [searchText, setSearchText] = useState<string>("")
  let [loading, setLoading] = useState<boolean>(false)
  let [items, setItems] = useState<NewsItem[]>([])
  let [pagination, setPagination] = useState({
    pn: 1,
    ps: 10,
  })
  let [isInit, setInit] = useState(true)
  const handleSearch = () => {
    setItems([]) // 清空列表
    appendData()
  }

  const appendData = () => {
    setLoading(true)
    SearchService.search(searchText, pagination)
      .then((res) => {
        setItems(items.concat(res.data.list))
        setLoading(false)
        message.success(`${res.data.list.length} more items loaded!`)
      })
      .catch((e) => {
        console.log(e)
        setLoading(false)
      })
  }

  useEffect(() => {
    if (isInit) {
      setInit(false)
    }else{
      appendData()
    }
    
  }, [pagination])

  const handleScroll = (e: React.UIEvent<HTMLElement, UIEvent>) => {
    if (e.currentTarget.scrollHeight <= e.currentTarget.scrollTop + e.currentTarget.clientHeight) {
      setPagination({pn: pagination.pn+1,ps: 10})
    }
  };

  return (
    <>
      <SearchInput
        onSearch={handleSearch}
        onChange={(event: any) => setSearchText(event.target.value)}
        loading={loading}
      />
      <NewsList
        data={items} 
        onScroll={handleScroll}
      />
    </>
  )
}

export default Search
