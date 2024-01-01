import { motion } from 'framer-motion';

function Loading() {
  return (
    <motion.div
      className="h-screen flex justify-center items-center"
      initial={ { opacity: 0 } }
      animate={ { opacity: 1 } }
      exit={ { opacity: 0, transition: { duration: 0.1 } } }
    >
      <div
        className="h-6 w-6 border-4 rounded-full mx-auto
     border-t-white border-l-white border-r-orange border-b-orange
    animate-spin"
      />
    </motion.div>
  );
}

export default Loading;
