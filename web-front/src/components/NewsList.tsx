import { List } from 'antd'
import VirtualList from 'rc-virtual-list'
import React from 'react'
import { RouteComponentProps } from '@reach/router'
import "../app.css"

interface NewsItem {
  news_url?: string
  title?: string
  rank?: number
  author?: string
  abstract?: string
  publish_time?: string
  is_hot?: boolean
  img_url?: string
  list_url?: string
  page_url?: string
  data_source?: string
}

interface Props extends React.PropsWithChildren, RouteComponentProps {
  onScroll?: (React.MouseEventHandler<HTMLButtonElement> | undefined)
  data: NewsItem[]
}

const NewsList: React.FC<Props> = (props) => {
  const ContainerHeight = 1000
  return (
    <List  className="search-list" bordered={false} pagination={false } split={true}>
      <VirtualList
        data={props.data}
        height={ContainerHeight}
        itemKey="news_url"
        onScroll={props.onScroll}
      >
        {(item: NewsItem) => (
          <List.Item key={item.news_url} className="item">
            <div className='container1'>
              <div className='news_title'>
                <a href={item.news_url} target="_blank" rel="noreferrer">{item.title}</a>
              </div>
              <div className='abstract'>
              <p>{item.abstract}</p>
              </div>
              <div className='tag'>
              排名:&nbsp;&nbsp;
                <em>{item.rank}</em>
              </div>
              <div className='tag'>
              发布时间:&nbsp;&nbsp;
                <em>{item.publish_time}</em>
              </div>
              <div className='tag'>
                <a href={item.page_url} target="_blank" rel="noreferrer">
                  来源:&nbsp;&nbsp;
                    <em>{item.data_source}</em>
                </a>
              </div>
            </div>
            
            <div className='container2'>
              <a href={item.news_url} target="_blank" rel="noreferrer">
              <img src={item.img_url} alt={"新闻图片不存在"} loading="lazy">    
              </img>
              </a>
            </div>
          </List.Item>
        )}
      </VirtualList>
    </List>
  );
};

export default NewsList;
export type {NewsItem};