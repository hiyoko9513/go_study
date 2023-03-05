package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/role"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// アノテーションを利用して、テーブル名を指定する1例
		entsql.Annotation{Table: "Users"},
		// カラムにコメントをデフォルトオンにする
		//entsql.WithComments(true),
		//schema.Comment("スキーマと生成されたコードに表示されるコメント"),
	}
}

type MyName string

// Fields of the User.
// IDフィールドについては特殊なので下記URLから確認
// https://entgo.io/ja/docs/schema-fields#id%E3%83%95%E3%82%A3%E3%83%BC%E3%83%AB%E3%83%89
// 組み込まれているバリデーター一覧
// https://entgo.io/ja/docs/schema-fields#%E7%B5%84%E3%81%BF%E8%BE%BC%E3%81%BF%E3%83%90%E3%83%AA%E3%83%87%E3%83%BC%E3%82%BF%E3%83%BC
func (User) Fields() []ent.Field {
	// サポートされている型 //bool //string //time.Time //UUID //[]byte (SQL only). //JSON (SQL only). //Enum (SQL only). //Other (SQLのみ)
	return []ent.Field{
		field.UUID("ulid", ulid.ULID{}).
			Default(ulid.New).
			StorageKey("id_ulid"). // todo もうちょっと詳しく確認（カラム名？？？）
			Immutable().
			Unique().
			Comment("こめんと出来るよ\n2行目"),
		// IDを独自に生成したい場合、関数を利用したい場合
		field.Int64("id_unix").
			DefaultFunc(func() int64 {
				return time.Now().Unix() << 8
			}),
		// uuid
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			Default("hiyoko"). // デフォルト設定
			Unique().
			Match(regexp.MustCompile("[a-zA-Z_]+$")). // カスタムバリデーター
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("グループは必ず大文字で始まる必要があります")
				}
				return nil
			}).
			Validate(MaxRuneCount(10)). // 使いまわしバリデーター
			Annotations(                // カラムコメントを明示的に無効にする
				entsql.WithComments(false),
			),
		// nilableとの違いについて
		field.String("required_name"),
		field.String("optional_name").
			Optional(),
		field.String("nillable_name").
			Optional().
			Nillable(),
		/////////////////////////
		field.String("custom_name").
			StructTag(`custom:"custom_name"`), // 生成する構造体にタグを追加出来る（json:nameのようにcustom:"custom_nameが追加される）
		// GOの型の利用
		field.String("go_name").
			Optional().
			GoType(MyName("")),
		field.String("value_scanner_name").
			Optional().
			GoType(&sql.NullString{}), // ValueScanner を実装している型
		field.String("password").
			Sensitive(), // センシティブ設定
		field.Enum("role").
			// string に変換可能な型
			GoType(role.Role("")).
			Values(role.Role.Values("")...), // enumのvalueを設定できる
		field.Int("age").
			Positive().
			Optional(), // entity作成時、任意項目に出来る(defaultでは必須項目)
		field.Float("rank").
			Optional(),
		field.Bool("active").
			Default(false),
		field.JSON("url", &url.URL{}).
			Optional(),
		field.JSON("strings", []string{}).
			Optional(),
		field.Enum("state").
			Values("on", "off").
			Optional(),
		field.JSON("dirs", []http.Dir{}).
			Default([]http.Dir{"/tmp"}),
		// DBの型へ変換
		// 下記の例: MySQLの場合はfloat64のフィールドをdouble型の列としてデータベースに作成します。
		field.Float("account_balance").
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)", // MySQL用の変換を上書き
				dialect.Postgres: "numeric",      // Postgres用の変換を上書き
			}),
		field.Time("created_at").
			Default(time.Now),
		// CURRENT_TIMESTAMPをデフォルト値とする
		// 新しいフィールドを追加
		field.Time("created_at_custom").
			Default(time.Now).
			Annotations(
				entsql.Default("CURRENT_TIMESTAMP"),
			).
			Immutable(), // entity生成時のみ設定出来るもの
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes インデックスの設定(詳しくは下記のURL)
// https://entgo.io/ja/docs/schema-indexes/
func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 非一意インデックス
		index.Fields("field1", "field2"),
		// 一意インデックス
		index.Fields("first_name", "last_name").
			Unique(),
	}
}

// MaxRuneCount unicode/utf8 パッケージを使って rune 単位での長さを検証します。
func MaxRuneCount(maxLen int) func(s string) error {
	return func(s string) error {
		if utf8.RuneCountInString(s) > maxLen {
			return errors.New("この文字列は長すぎます")
		}
		return nil
	}
}
