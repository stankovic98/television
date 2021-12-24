// export const getAllVids = async () => {
//   let videos = await window.backend.GetAllVideos();
//   return videos;
// };

export const getPaginatedVideos = async (from, to) => {
  let videos = await window.backend.getVideos(from, to);
  return videos;
};

export const getNumOfVideos = async () => {
  let num = await window.backend.getNumOfVids();
  return num;
};
