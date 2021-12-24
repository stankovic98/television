import React, { useEffect, useState } from "react";
import { getNumOfVideos, getPaginatedVideos } from "../services/backend";
import VideoBox from "../components/VideoBox";

const paginationStep = 12;

const Feed = () => {
  const [videos, setVideos] = useState([]);
  const [selectedVideos, setSelectedVideos] = useState(0);
  const [numOfAllVids, setNumOfAllVids] = useState(0);
  const [isLoading, setIsLoading] = useState(true);

  const initData = async () => {
    let res = await getPaginatedVideos(
      selectedVideos,
      selectedVideos + paginationStep
    );
    setVideos(res);
    setIsLoading(false);
    let num = await getNumOfVideos();
    setNumOfAllVids(num);
  };

  const setPagination = () => {
    let paginationBlock = [];
    // paginationBlock.push(<a href="#">&laquo;</a>);
    for (let i = 0; i < numOfAllVids; i += paginationStep) {
      if (i === selectedVideos) {
        paginationBlock.push(
          <a href="#" className="selected" key={i}>
            {i / paginationStep + 1}
          </a>
        );
      } else {
        paginationBlock.push(
          <a href="#" onClick={(e) => getPage(e, i)}>
            {i / paginationStep + 1}
          </a>
        );
      }
    }
    // paginationBlock.push(<a href="#">&raquo;</a>);
    return paginationBlock;
  };

  const getPage = async (e, i) => {
    e.preventDefault();
    setIsLoading(true);
    setSelectedVideos(i);
    let res = await getPaginatedVideos(i, i + paginationStep);
    setVideos(res);
    setIsLoading(false);
  };

  useEffect(() => {
    initData();
  }, []);

  return (
    <div>
      <h1>Become Virtuous</h1>
      <div className="container">
        {isLoading && <div className="lds-dual-ring"></div>}
        {videos && videos.map((vid) => <VideoBox video={vid} key={vid.ID} />)}
        {/* definetly needs pagination */}
        <div className="pagination">{setPagination()}</div>
      </div>
    </div>
  );
};

export default Feed;
