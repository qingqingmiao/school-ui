-- 1. 数据库-tb_school_list

CREATE TABLE "public"."tb_school_list"(
	"school_id" SERIAL PRIMARY KEY,
	"school_code" INT,
	"school_name" VARCHAR(50) NOT NULL,
	"school_province" VARCHAR(50),
	"school_city" VARCHAR(50),
	"school_department" VARCHAR(50),
	"school_level" VARCHAR(50),
	"school_type" VARCHAR(50),
	"school_website" VARCHAR(255)
)


-- 2. 数据库-tb_school_province

CREATE TABLE "public"."tb_school_province"(
	"school_id" INT PRIMARY KEY,
	"province_id" INT,
	"province_name" VARCHAR(50)
)


-- 3. 数据库-tb_school_city

CREATE TABLE "public"."tb_school_city"(
	"school_id" INT PRIMARY KEY,
	"city_id" INT,
	"city_name" VARCHAR(50)
)


-- 4. 数据库-tb_school_type

CREATE TABLE "public"."tb_school_type"(
	"school_id" INT PRIMARY KEY,
	"is_dstype" BOOLEAN NOT NULL,
	"is_first_class" BOOLEAN NOT NULL,
	"is_985" BOOLEAN NOT NULL,
	"is_211" BOOLEAN NOT NULL
)