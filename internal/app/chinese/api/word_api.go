package api

import (
	"github.com/maldan/gam-app-chinese/internal/app/chinese/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-rapi/rapi_core"
	"strings"
)

type WordApi struct {
}

// GetIndex Get word
func (r WordApi) GetIndex(args core.Word) core.Word {
	out := core.Word{}
	err := cmhp_file.ReadJSON(core.DataDir+"/word/"+args.Name+".json", &out)
	rapi_core.FatalIfError(err)

	if out.Translate.Noun == nil {
		out.Translate.Noun = make([]string, 0)
	}
	if out.Translate.Verb == nil {
		out.Translate.Verb = make([]string, 0)
	}
	if out.Translate.Adjective == nil {
		out.Translate.Adjective = make([]string, 0)
	}
	if out.Translate.Adverb == nil {
		out.Translate.Adverb = make([]string, 0)
	}
	if out.Translate.Particle == nil {
		out.Translate.Particle = make([]string, 0)
	}
	if out.Translate.Preposition == nil {
		out.Translate.Preposition = make([]string, 0)
	}

	return out
}

// GetList Get word list
func (r WordApi) GetList() []core.Word {
	out := make([]core.Word, 0)
	l, _ := cmhp_file.List(core.DataDir + "/word")
	for _, f := range l {
		word := r.GetIndex(core.Word{Name: strings.Replace(f.Name(), ".json", "", 1)})
		out = append(out, word)
	}
	return out
}

// PostIndex Save word
func (r WordApi) PostIndex(args core.Word) {
	err := cmhp_file.Write(core.DataDir+"/word/"+args.Name+".json", &args)
	rapi_core.FatalIfError(err)
}

// PatchIndex Save word
func (r WordApi) PatchIndex(args core.Word) {
	err := cmhp_file.Write(core.DataDir+"/word/"+args.Name+".json", &args)
	rapi_core.FatalIfError(err)
}

// DeleteIndex Delete word
func (r WordApi) DeleteIndex(args core.Word) {
	err := cmhp_file.Delete(core.DataDir + "/word/" + args.Name + ".json")
	rapi_core.FatalIfError(err)
}
