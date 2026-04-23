-- phpMyAdmin SQL Dump
-- version 5.2.1deb3
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Apr 23, 2026 at 11:17 AM
-- Server version: 10.11.14-MariaDB-0ubuntu0.24.04.1
-- PHP Version: 8.3.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_study`
--

-- --------------------------------------------------------

--
-- Table structure for table `activity_logs`
--

CREATE TABLE `activity_logs` (
  `id` int(11) NOT NULL,
  `ip` varchar(45) NOT NULL,
  `aktivitas` text NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `activity_logs`
--

INSERT INTO `activity_logs` (`id`, `ip`, `aktivitas`, `created_at`) VALUES
(1, '::1', 'Menambahkan data siswa Pak Ahmad', '2026-04-22 12:39:37'),
(2, '::1', 'Mengambil list data guru', '2026-04-22 13:27:46'),
(3, '::1', 'Mengambil list data kelas', '2026-04-22 13:51:05'),
(4, '::1', 'Menambah data siswa', '2026-04-23 01:06:04'),
(5, '::1', 'Menambah data siswa', '2026-04-23 01:06:19'),
(6, '::1', 'Menambah data siswa', '2026-04-23 01:06:41'),
(7, '::1', 'Menambah data siswa', '2026-04-23 01:06:59'),
(8, '::1', 'Menambah data siswa', '2026-04-23 01:07:10'),
(9, '::1', 'Mengambil list data siswa', '2026-04-23 01:07:22'),
(10, '::1', 'Mengambil list data siswa', '2026-04-23 01:08:23'),
(11, '::1', 'Mengambil list data siswa', '2026-04-23 01:10:17'),
(12, '::1', 'Mengambil list data siswa', '2026-04-23 01:13:10'),
(13, '::1', 'Mengambil list data siswa', '2026-04-23 01:13:21'),
(14, '::1', 'Mengambil list data siswa', '2026-04-23 01:14:53'),
(15, '::1', 'Mengambil list data siswa', '2026-04-23 01:15:39'),
(16, '::1', 'Mengambil list data siswa', '2026-04-23 01:15:48'),
(17, '::1', 'Mengambil list data siswa', '2026-04-23 01:16:05'),
(18, '::1', 'Mengambil list data siswa', '2026-04-23 01:16:38'),
(19, '::1', 'Mengambil list data siswa', '2026-04-23 01:17:13'),
(20, '::1', 'Mengambil list data siswa', '2026-04-23 01:17:27'),
(21, '::1', 'Mengambil list data siswa', '2026-04-23 01:17:41'),
(22, '::1', 'Mengambil list data siswa', '2026-04-23 01:18:05'),
(23, '::1', 'Menambah data siswa', '2026-04-23 01:40:47'),
(24, '::1', 'Mengambil list data siswa', '2026-04-23 01:41:17'),
(25, '::1', 'Mengambil list data siswa', '2026-04-23 01:41:43'),
(26, '::1', 'Mengambil list data siswa', '2026-04-23 01:42:12'),
(27, '::1', 'Mengambil list data siswa', '2026-04-23 01:42:25'),
(28, '::1', 'Mengambil list data siswa', '2026-04-23 01:42:32'),
(29, '::1', 'Mengambil data siswa', '2026-04-23 01:42:52'),
(30, '::1', 'Mengubah data siswa', '2026-04-23 01:44:05'),
(31, '::1', 'Menghapus data siswa', '2026-04-23 01:44:19'),
(32, '::1', 'Mengambil list detail data siswa', '2026-04-23 01:46:13'),
(33, '::1', 'Mengambil list data pelajaran', '2026-04-23 01:46:35'),
(34, '::1', 'Mengambil list detail data siswa', '2026-04-23 01:46:58'),
(35, '::1', 'Mengambil list data kelas', '2026-04-23 01:48:27'),
(36, '::1', 'Mengambil list data kelas', '2026-04-23 01:48:56'),
(37, '::1', 'Mengambil list data kelas', '2026-04-23 01:49:12'),
(38, '::1', 'Mengambil list data kelas', '2026-04-23 01:51:52'),
(39, '::1', 'Mengambil list rata-rata nilai siswa', '2026-04-23 01:52:21'),
(40, '::1', 'Menambahkan data guru Pak farid', '2026-04-23 01:52:57'),
(41, '::1', 'Mengambil list data guru', '2026-04-23 01:53:07'),
(42, '::1', 'Mengambil data guru', '2026-04-23 01:59:17'),
(43, '::1', 'Menghapus data guru', '2026-04-23 01:59:39'),
(44, '::1', 'Mengambil list data guru', '2026-04-23 02:50:56'),
(45, '::1', 'Mengambil data guru', '2026-04-23 02:51:56'),
(46, '::1', 'Mengambil list data guru', '2026-04-23 02:53:52'),
(47, '::1', 'Mengambil list data siswa', '2026-04-23 02:58:53'),
(48, '::1', 'Mengambil list data siswa', '2026-04-23 02:59:03'),
(49, '::1', 'Mengambil list data siswa', '2026-04-23 03:01:25'),
(50, '::1', 'Mengambil list data guru', '2026-04-23 03:05:55'),
(51, '::1', 'Mengambil list data guru', '2026-04-23 03:06:11'),
(52, '::1', 'Mengambil list data activity log', '2026-04-23 03:40:21'),
(53, '::1', 'Mengambil data activity log', '2026-04-23 03:42:02'),
(54, '::1', 'Menambah data siswa', '2026-04-23 03:45:58'),
(55, '::1', 'Mengambil list data siswa', '2026-04-23 03:52:02'),
(56, '::1', 'Mengambil list data siswa', '2026-04-23 03:52:32'),
(57, '::1', 'Mengambil list data siswa', '2026-04-23 03:52:42'),
(58, '::1', 'Mengambil list data siswa', '2026-04-23 03:52:48'),
(59, '::1', 'Mengambil list data siswa', '2026-04-23 03:54:03'),
(60, '::1', 'Mengambil list data siswa', '2026-04-23 03:54:41'),
(61, '::1', 'Mengambil list data siswa', '2026-04-23 03:54:52'),
(62, '::1', 'Mengambil list data siswa', '2026-04-23 03:55:10'),
(63, '::1', 'Mengambil list data siswa', '2026-04-23 03:55:25'),
(64, '::1', 'Mengambil data siswa', '2026-04-23 03:55:48'),
(65, '::1', 'Mengambil data siswa', '2026-04-23 03:56:09'),
(66, '::1', 'Mengambil data siswa', '2026-04-23 03:56:30'),
(67, '::1', 'Mengubah data siswa', '2026-04-23 03:59:37'),
(68, '::1', 'Mengubah data siswa', '2026-04-23 04:00:00'),
(69, '::1', 'Mengubah data siswa', '2026-04-23 04:00:14'),
(70, '::1', 'Mengubah data siswa', '2026-04-23 04:01:08'),
(71, '::1', 'Mengubah data siswa', '2026-04-23 04:18:56'),
(72, '::1', 'Mengubah data siswa', '2026-04-23 04:21:12'),
(73, '::1', 'Menghapus data siswa', '2026-04-23 04:27:51'),
(74, '::1', 'Mengambil list detail data siswa', '2026-04-23 04:28:20'),
(75, '::1', 'Mengambil list rata-rata nilai siswa', '2026-04-23 04:28:29'),
(76, '::1', 'Menambahkan data guru Pak salman', '2026-04-23 04:29:57'),
(77, '::1', 'Mengambil list data guru', '2026-04-23 04:30:20'),
(78, '::1', 'Mengambil list data guru', '2026-04-23 04:31:03'),
(79, '::1', 'Mengambil list data guru', '2026-04-23 04:31:11'),
(80, '::1', 'Mengambil list data guru', '2026-04-23 04:31:41'),
(81, '::1', 'Mengambil data guru', '2026-04-23 04:31:56'),
(82, '::1', 'Menghapus data guru', '2026-04-23 04:32:53'),
(83, '::1', 'Mengubah data guru', '2026-04-23 04:33:23'),
(84, '::1', 'Mengubah data guru', '2026-04-23 04:33:41'),
(85, '::1', 'Mengubah data guru', '2026-04-23 04:33:54'),
(86, '::1', 'Menambah data pelajaran', '2026-04-23 04:35:10'),
(87, '::1', 'Mengambil list data pelajaran', '2026-04-23 04:35:22'),
(88, '::1', 'Mengambil data pelajaran', '2026-04-23 04:35:34'),
(89, '::1', 'Mengubah data pelajaran', '2026-04-23 04:35:48'),
(90, '::1', 'Menghapus data pelajaran', '2026-04-23 04:36:39'),
(91, '::1', 'Menambah nilai siswa', '2026-04-23 04:37:25'),
(92, '::1', 'Mengambil list nilai siswa', '2026-04-23 04:37:38'),
(93, '::1', 'Mengambil nilai siswa', '2026-04-23 04:38:35'),
(94, '::1', 'Mengubah nilai siswa', '2026-04-23 04:40:13'),
(95, '::1', 'Menghapus nilai siswa', '2026-04-23 04:40:49'),
(96, '::1', 'Menambah data kelas', '2026-04-23 04:41:47'),
(97, '::1', 'Mengambil list data kelas', '2026-04-23 04:41:57'),
(98, '::1', 'Mengambil data kelas', '2026-04-23 04:42:04'),
(99, '::1', 'Mengubah data kelas', '2026-04-23 04:42:20'),
(100, '::1', 'Menghapus data kelas', '2026-04-23 04:42:53'),
(101, '::1', 'Mengambil list data activity log', '2026-04-23 04:43:07'),
(102, '::1', 'Mengambil list data activity log', '2026-04-23 04:43:39'),
(103, '::1', 'Mengambil list data activity log', '2026-04-23 04:43:52'),
(104, '::1', 'Mengambil data activity log', '2026-04-23 04:44:03'),
(105, '::1', 'Mengambil data activity log', '2026-04-23 04:44:20'),
(106, '::1', 'Mengubah data pelajaran', '2026-04-23 06:11:22'),
(107, '::1', 'Mengubah nilai siswa', '2026-04-23 06:19:48'),
(108, '::1', 'Mengubah nilai siswa', '2026-04-23 06:27:16'),
(109, '::1', 'Mengambil list data siswa', '2026-04-23 09:47:00'),
(110, '::1', 'Mengambil list data siswa', '2026-04-23 09:47:26'),
(111, '::1', 'Mengambil list data siswa', '2026-04-23 09:47:39'),
(112, '::1', 'Mengambil list data activity log', '2026-04-23 09:49:25'),
(113, '::1', 'Mengambil list data activity log', '2026-04-23 09:51:32'),
(114, '::1', 'Mengambil list data activity log', '2026-04-23 09:52:42'),
(115, '::1', 'Mengambil list data siswa', '2026-04-23 09:56:27'),
(116, '::1', 'Mengambil list data activity log', '2026-04-23 09:58:13'),
(117, '::1', 'Mengambil list data activity log', '2026-04-23 10:00:34'),
(118, '::1', 'Mengambil list data activity log', '2026-04-23 10:00:49'),
(119, '::1', 'Mengambil list data activity log', '2026-04-23 10:03:45'),
(120, '::1', 'Mengambil list data activity log', '2026-04-23 10:07:42'),
(121, '::1', 'Mengambil list data activity log', '2026-04-23 10:10:34'),
(122, '::1', 'Mengambil list data activity log', '2026-04-23 10:13:39'),
(123, '::1', 'Mengambil list data activity log', '2026-04-23 10:15:15'),
(124, '::1', 'Mengambil list data siswa', '2026-04-23 10:16:12'),
(125, '::1', 'Mengambil list data activity log', '2026-04-23 10:20:53'),
(126, '::1', 'Mengambil list data kelas', '2026-04-23 10:21:52'),
(127, '::1', 'Mengambil list data activity log', '2026-04-23 10:22:01'),
(128, '::1', 'Mengambil list data siswa', '2026-04-23 10:57:26'),
(129, '::1', 'Mengambil list data log', '2026-04-23 10:57:31'),
(130, '::1', 'Mengambil list data log', '2026-04-23 10:57:56'),
(131, '::1', 'Mengambil list data siswa', '2026-04-23 11:12:23'),
(132, '::1', 'Mengambil list data log', '2026-04-23 11:13:31'),
(133, '::1', 'Mengambil list data siswa', '2026-04-23 11:14:04'),
(134, '::1', 'Mengambil list data log', '2026-04-23 11:14:19');

-- --------------------------------------------------------

--
-- Table structure for table `classes`
--

CREATE TABLE `classes` (
  `id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `teacher_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `classes`
--

INSERT INTO `classes` (`id`, `name`, `teacher_id`) VALUES
(1, 'Kelas 10-A', 1),
(2, 'Class 10-B', 8),
(3, 'Kelas 9', 2);

-- --------------------------------------------------------

--
-- Table structure for table `grades`
--

CREATE TABLE `grades` (
  `id` int(11) NOT NULL,
  `student_id` int(11) DEFAULT NULL,
  `subject_id` int(11) DEFAULT NULL,
  `score` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `grades`
--

INSERT INTO `grades` (`id`, `student_id`, `subject_id`, `score`) VALUES
(1, 1, 2, 97),
(3, 2, 4, 98),
(4, 3, 2, 88),
(5, 3, 3, 82),
(6, 6, 4, 75),
(7, 8, 3, 75);

-- --------------------------------------------------------

--
-- Table structure for table `schema_migrations`
--

CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `schema_migrations`
--

INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
(6, 0);

-- --------------------------------------------------------

--
-- Table structure for table `students`
--

CREATE TABLE `students` (
  `id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `class_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `students`
--

INSERT INTO `students` (`id`, `name`, `email`, `class_id`) VALUES
(1, 'Andi Wijaya', 'andi.wijaya@student.id', 1),
(2, 'Bunga', 'bunga@mail.com', 1),
(3, 'Candra Gupta', 'candra.gupta@student.id', 2),
(4, 'Antin', 'antin@student.id', 1),
(6, 'Lisa', 'lisa@student.id', 1),
(8, 'moana', 'moana@student.id', 1),
(9, 'dims', 'dimas@mail.com', 3),
(11, 'ahmad', 'ahmad@mail.com', 3),
(21, 'burhan', 'burhan@mail.com', 2),
(22, 'hakim', 'hakim@mail.com', 2),
(23, 'danish', 'danish@mail.com', 3),
(24, 'brian', 'brian@mail.com', 3),
(25, 'veda', 'veda@mail.com', 3);

-- --------------------------------------------------------

--
-- Table structure for table `subjects`
--

CREATE TABLE `subjects` (
  `id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `teacher_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `subjects`
--

INSERT INTO `subjects` (`id`, `name`, `teacher_id`) VALUES
(1, 'Matematika', 1),
(2, 'Bahasa Inggris', 8),
(3, 'Fisika', 1),
(4, 'IPS', 2),
(7, 'Sejarah', 2);

-- --------------------------------------------------------

--
-- Table structure for table `teachers`
--

CREATE TABLE `teachers` (
  `id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `teachers`
--

INSERT INTO `teachers` (`id`, `name`, `email`) VALUES
(1, 'Budi Santoso', 'budi.santoso@sekolah.id'),
(2, 'Bu Siti Aminah', 'bu.siti.aminah@sekolah.id'),
(8, 'Pak salman', 'pak.salman@sekolah.id');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `activity_logs`
--
ALTER TABLE `activity_logs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `classes`
--
ALTER TABLE `classes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `teacher_id` (`teacher_id`);

--
-- Indexes for table `grades`
--
ALTER TABLE `grades`
  ADD PRIMARY KEY (`id`),
  ADD KEY `student_id` (`student_id`),
  ADD KEY `subject_id` (`subject_id`);

--
-- Indexes for table `schema_migrations`
--
ALTER TABLE `schema_migrations`
  ADD PRIMARY KEY (`version`);

--
-- Indexes for table `students`
--
ALTER TABLE `students`
  ADD PRIMARY KEY (`id`),
  ADD KEY `class_id` (`class_id`);

--
-- Indexes for table `subjects`
--
ALTER TABLE `subjects`
  ADD PRIMARY KEY (`id`),
  ADD KEY `teacher_id` (`teacher_id`);

--
-- Indexes for table `teachers`
--
ALTER TABLE `teachers`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `activity_logs`
--
ALTER TABLE `activity_logs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=135;

--
-- AUTO_INCREMENT for table `classes`
--
ALTER TABLE `classes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `grades`
--
ALTER TABLE `grades`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `students`
--
ALTER TABLE `students`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- AUTO_INCREMENT for table `subjects`
--
ALTER TABLE `subjects`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `teachers`
--
ALTER TABLE `teachers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `classes`
--
ALTER TABLE `classes`
  ADD CONSTRAINT `classes_ibfk_1` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`);

--
-- Constraints for table `grades`
--
ALTER TABLE `grades`
  ADD CONSTRAINT `grades_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`),
  ADD CONSTRAINT `grades_ibfk_2` FOREIGN KEY (`subject_id`) REFERENCES `subjects` (`id`);

--
-- Constraints for table `students`
--
ALTER TABLE `students`
  ADD CONSTRAINT `students_ibfk_1` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`);

--
-- Constraints for table `subjects`
--
ALTER TABLE `subjects`
  ADD CONSTRAINT `subjects_ibfk_1` FOREIGN KEY (`teacher_id`) REFERENCES `teachers` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
