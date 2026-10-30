// generated code - do not edit
package aixmfeatures

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.A_extensions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension", instances, (*A_extension)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_1", instances, (*A_extension_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_10s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_10", instances, (*A_extension_10)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_100s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_100", instances, (*A_extension_100)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_101s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_101", instances, (*A_extension_101)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_102s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_102", instances, (*A_extension_102)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_103s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_103", instances, (*A_extension_103)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_104s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_104", instances, (*A_extension_104)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_105s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_105", instances, (*A_extension_105)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_106s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_106", instances, (*A_extension_106)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_107s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_107", instances, (*A_extension_107)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_108s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_108", instances, (*A_extension_108)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_109s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_109", instances, (*A_extension_109)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_11s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_11", instances, (*A_extension_11)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_110s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_110", instances, (*A_extension_110)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_111s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_111", instances, (*A_extension_111)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_112s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_112", instances, (*A_extension_112)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_113s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_113", instances, (*A_extension_113)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_114s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_114", instances, (*A_extension_114)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_115s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_115", instances, (*A_extension_115)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_116s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_116", instances, (*A_extension_116)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_117s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_117", instances, (*A_extension_117)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_118s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_118", instances, (*A_extension_118)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_119s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_119", instances, (*A_extension_119)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_12s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_12", instances, (*A_extension_12)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_120s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_120", instances, (*A_extension_120)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_121s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_121", instances, (*A_extension_121)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_122s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_122", instances, (*A_extension_122)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_123s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_123", instances, (*A_extension_123)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_124s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_124", instances, (*A_extension_124)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_125s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_125", instances, (*A_extension_125)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_126s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_126", instances, (*A_extension_126)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_127s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_127", instances, (*A_extension_127)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_128s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_128", instances, (*A_extension_128)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_129s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_129", instances, (*A_extension_129)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_13s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_13", instances, (*A_extension_13)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_130s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_130", instances, (*A_extension_130)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_131s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_131", instances, (*A_extension_131)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_132s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_132", instances, (*A_extension_132)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_133s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_133", instances, (*A_extension_133)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_134s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_134", instances, (*A_extension_134)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_135s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_135", instances, (*A_extension_135)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_136s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_136", instances, (*A_extension_136)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_137s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_137", instances, (*A_extension_137)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_138s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_138", instances, (*A_extension_138)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_139s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_139", instances, (*A_extension_139)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_14s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_14", instances, (*A_extension_14)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_140s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_140", instances, (*A_extension_140)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_141s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_141", instances, (*A_extension_141)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_142s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_142", instances, (*A_extension_142)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_143s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_143", instances, (*A_extension_143)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_144s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_144", instances, (*A_extension_144)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_145s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_145", instances, (*A_extension_145)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_146s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_146", instances, (*A_extension_146)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_147s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_147", instances, (*A_extension_147)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_148s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_148", instances, (*A_extension_148)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_149s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_149", instances, (*A_extension_149)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_15s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_15", instances, (*A_extension_15)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_150s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_150", instances, (*A_extension_150)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_151s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_151", instances, (*A_extension_151)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_152s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_152", instances, (*A_extension_152)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_153s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_153", instances, (*A_extension_153)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_154s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_154", instances, (*A_extension_154)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_155s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_155", instances, (*A_extension_155)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_156s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_156", instances, (*A_extension_156)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_157s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_157", instances, (*A_extension_157)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_158s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_158", instances, (*A_extension_158)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_159s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_159", instances, (*A_extension_159)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_16s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_16", instances, (*A_extension_16)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_160s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_160", instances, (*A_extension_160)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_161s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_161", instances, (*A_extension_161)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_162s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_162", instances, (*A_extension_162)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_163s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_163", instances, (*A_extension_163)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_164s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_164", instances, (*A_extension_164)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_165s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_165", instances, (*A_extension_165)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_166s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_166", instances, (*A_extension_166)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_167s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_167", instances, (*A_extension_167)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_168s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_168", instances, (*A_extension_168)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_169s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_169", instances, (*A_extension_169)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_17s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_17", instances, (*A_extension_17)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_170s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_170", instances, (*A_extension_170)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_171s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_171", instances, (*A_extension_171)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_172s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_172", instances, (*A_extension_172)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_173s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_173", instances, (*A_extension_173)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_174s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_174", instances, (*A_extension_174)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_175s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_175", instances, (*A_extension_175)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_176s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_176", instances, (*A_extension_176)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_177s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_177", instances, (*A_extension_177)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_178s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_178", instances, (*A_extension_178)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_179s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_179", instances, (*A_extension_179)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_18s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_18", instances, (*A_extension_18)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_180s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_180", instances, (*A_extension_180)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_181s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_181", instances, (*A_extension_181)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_182s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_182", instances, (*A_extension_182)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_183s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_183", instances, (*A_extension_183)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_184s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_184", instances, (*A_extension_184)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_185s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_185", instances, (*A_extension_185)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_186s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_186", instances, (*A_extension_186)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_187s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_187", instances, (*A_extension_187)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_188s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_188", instances, (*A_extension_188)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_189s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_189", instances, (*A_extension_189)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_19s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_19", instances, (*A_extension_19)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_190s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_190", instances, (*A_extension_190)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_191s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_191", instances, (*A_extension_191)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_192s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_192", instances, (*A_extension_192)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_193s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_193", instances, (*A_extension_193)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_194s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_194", instances, (*A_extension_194)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_195s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_195", instances, (*A_extension_195)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_196s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_196", instances, (*A_extension_196)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_197s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_197", instances, (*A_extension_197)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_198s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_198", instances, (*A_extension_198)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_199s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_199", instances, (*A_extension_199)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_2s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_2", instances, (*A_extension_2)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_20s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_20", instances, (*A_extension_20)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_200s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_200", instances, (*A_extension_200)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_201s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_201", instances, (*A_extension_201)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_202s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_202", instances, (*A_extension_202)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_203s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_203", instances, (*A_extension_203)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_204s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_204", instances, (*A_extension_204)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_205s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_205", instances, (*A_extension_205)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_206s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_206", instances, (*A_extension_206)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_207s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_207", instances, (*A_extension_207)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_208s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_208", instances, (*A_extension_208)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_209s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_209", instances, (*A_extension_209)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_21s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_21", instances, (*A_extension_21)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_210s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_210", instances, (*A_extension_210)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_211s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_211", instances, (*A_extension_211)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_212s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_212", instances, (*A_extension_212)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_213s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_213", instances, (*A_extension_213)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_214s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_214", instances, (*A_extension_214)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_215s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_215", instances, (*A_extension_215)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_216s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_216", instances, (*A_extension_216)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_217s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_217", instances, (*A_extension_217)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_218s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_218", instances, (*A_extension_218)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_219s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_219", instances, (*A_extension_219)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_22s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_22", instances, (*A_extension_22)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_220s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_220", instances, (*A_extension_220)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_221s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_221", instances, (*A_extension_221)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_222s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_222", instances, (*A_extension_222)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_223s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_223", instances, (*A_extension_223)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_224s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_224", instances, (*A_extension_224)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_225s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_225", instances, (*A_extension_225)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_226s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_226", instances, (*A_extension_226)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_227s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_227", instances, (*A_extension_227)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_228s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_228", instances, (*A_extension_228)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_229s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_229", instances, (*A_extension_229)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_23s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_23", instances, (*A_extension_23)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_230s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_230", instances, (*A_extension_230)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_231s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_231", instances, (*A_extension_231)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_232s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_232", instances, (*A_extension_232)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_233s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_233", instances, (*A_extension_233)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_234s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_234", instances, (*A_extension_234)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_235s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_235", instances, (*A_extension_235)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_236s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_236", instances, (*A_extension_236)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_237s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_237", instances, (*A_extension_237)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_238s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_238", instances, (*A_extension_238)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_239s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_239", instances, (*A_extension_239)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_24s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_24", instances, (*A_extension_24)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_240s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_240", instances, (*A_extension_240)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_241s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_241", instances, (*A_extension_241)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_25s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_25", instances, (*A_extension_25)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_26s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_26", instances, (*A_extension_26)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_27s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_27", instances, (*A_extension_27)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_28s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_28", instances, (*A_extension_28)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_29s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_29", instances, (*A_extension_29)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_3s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_3", instances, (*A_extension_3)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_30s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_30", instances, (*A_extension_30)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_31s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_31", instances, (*A_extension_31)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_32s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_32", instances, (*A_extension_32)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_33s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_33", instances, (*A_extension_33)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_34s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_34", instances, (*A_extension_34)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_35s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_35", instances, (*A_extension_35)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_36s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_36", instances, (*A_extension_36)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_37s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_37", instances, (*A_extension_37)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_38s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_38", instances, (*A_extension_38)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_39s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_39", instances, (*A_extension_39)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_4s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_4", instances, (*A_extension_4)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_40s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_40", instances, (*A_extension_40)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_41s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_41", instances, (*A_extension_41)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_42s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_42", instances, (*A_extension_42)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_43s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_43", instances, (*A_extension_43)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_44s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_44", instances, (*A_extension_44)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_45s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_45", instances, (*A_extension_45)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_46s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_46", instances, (*A_extension_46)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_47s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_47", instances, (*A_extension_47)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_48s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_48", instances, (*A_extension_48)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_49s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_49", instances, (*A_extension_49)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_5s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_5", instances, (*A_extension_5)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_50s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_50", instances, (*A_extension_50)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_51s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_51", instances, (*A_extension_51)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_52s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_52", instances, (*A_extension_52)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_53s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_53", instances, (*A_extension_53)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_54s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_54", instances, (*A_extension_54)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_55s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_55", instances, (*A_extension_55)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_56s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_56", instances, (*A_extension_56)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_57s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_57", instances, (*A_extension_57)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_58s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_58", instances, (*A_extension_58)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_59s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_59", instances, (*A_extension_59)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_6s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_6", instances, (*A_extension_6)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_60s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_60", instances, (*A_extension_60)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_61s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_61", instances, (*A_extension_61)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_62s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_62", instances, (*A_extension_62)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_63s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_63", instances, (*A_extension_63)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_64s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_64", instances, (*A_extension_64)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_65s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_65", instances, (*A_extension_65)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_66s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_66", instances, (*A_extension_66)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_67s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_67", instances, (*A_extension_67)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_68s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_68", instances, (*A_extension_68)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_69s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_69", instances, (*A_extension_69)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_7s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_7", instances, (*A_extension_7)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_70s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_70", instances, (*A_extension_70)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_71s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_71", instances, (*A_extension_71)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_72s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_72", instances, (*A_extension_72)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_73s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_73", instances, (*A_extension_73)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_74s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_74", instances, (*A_extension_74)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_75s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_75", instances, (*A_extension_75)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_76s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_76", instances, (*A_extension_76)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_77s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_77", instances, (*A_extension_77)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_78s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_78", instances, (*A_extension_78)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_79s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_79", instances, (*A_extension_79)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_8s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_8", instances, (*A_extension_8)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_80s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_80", instances, (*A_extension_80)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_81s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_81", instances, (*A_extension_81)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_82s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_82", instances, (*A_extension_82)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_83s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_83", instances, (*A_extension_83)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_84s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_84", instances, (*A_extension_84)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_85s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_85", instances, (*A_extension_85)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_86s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_86", instances, (*A_extension_86)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_87s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_87", instances, (*A_extension_87)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_88s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_88", instances, (*A_extension_88)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_89s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_89", instances, (*A_extension_89)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_9s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_9", instances, (*A_extension_9)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_90s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_90", instances, (*A_extension_90)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_91s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_91", instances, (*A_extension_91)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_92s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_92", instances, (*A_extension_92)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_93s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_93", instances, (*A_extension_93)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_94s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_94", instances, (*A_extension_94)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_95s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_95", instances, (*A_extension_95)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_96s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_96", instances, (*A_extension_96)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_97s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_97", instances, (*A_extension_97)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_98s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_98", instances, (*A_extension_98)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_99s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_99", instances, (*A_extension_99)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractAirportGroundServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractAirportGroundServiceType", instances, (*AbstractAirportGroundServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractAirportHeliportProtectionAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractAirportHeliportProtectionAreaType", instances, (*AbstractAirportHeliportProtectionAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractApproachLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractApproachLegType", instances, (*AbstractApproachLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDirectFlightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDirectFlightType", instances, (*AbstractDirectFlightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGroundLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGroundLightSystemType", instances, (*AbstractGroundLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractMarkingType", instances, (*AbstractMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractNavaidEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractNavaidEquipmentType", instances, (*AbstractNavaidEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractNavigationSystemCheckpointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractNavigationSystemCheckpointType", instances, (*AbstractNavigationSystemCheckpointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractProcedureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractProcedureType", instances, (*AbstractProcedureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractPropertiesWithScheduleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractPropertiesWithScheduleType", instances, (*AbstractPropertiesWithScheduleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractRadarEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractRadarEquipmentType", instances, (*AbstractRadarEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSegmentLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSegmentLegType", instances, (*AbstractSegmentLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSegmentPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSegmentPointType", instances, (*AbstractSegmentPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractServiceType", instances, (*AbstractServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSurfaceContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSurfaceContaminationType", instances, (*AbstractSurfaceContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSurveillanceRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSurveillanceRadarType", instances, (*AbstractSurveillanceRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTrafficSeparationServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTrafficSeparationServiceType", instances, (*AbstractTrafficSeparationServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractUsageConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractUsageConditionType", instances, (*AbstractUsageConditionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingAnchorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingAnchorPropertyType", instances, (*AerialRefuellingAnchorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingAnchorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingAnchorType", instances, (*AerialRefuellingAnchorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingPointPropertyType", instances, (*AerialRefuellingPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingPointType", instances, (*AerialRefuellingPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingPropertyType", instances, (*AerialRefuellingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTimeSlicePropertyType", instances, (*AerialRefuellingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTimeSliceType", instances, (*AerialRefuellingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTrackPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTrackPropertyType", instances, (*AerialRefuellingTrackPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTrackTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingTrackType", instances, (*AerialRefuellingTrackType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AerialRefuellingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AerialRefuellingType", instances, (*AerialRefuellingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AeronauticalGroundLightPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AeronauticalGroundLightPropertyType", instances, (*AeronauticalGroundLightPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AeronauticalGroundLightTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AeronauticalGroundLightTimeSlicePropertyType", instances, (*AeronauticalGroundLightTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AeronauticalGroundLightTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AeronauticalGroundLightTimeSliceType", instances, (*AeronauticalGroundLightTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AeronauticalGroundLightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AeronauticalGroundLightType", instances, (*AeronauticalGroundLightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficControlServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficControlServicePropertyType", instances, (*AirTrafficControlServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficControlServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficControlServiceTimeSlicePropertyType", instances, (*AirTrafficControlServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficControlServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficControlServiceTimeSliceType", instances, (*AirTrafficControlServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficControlServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficControlServiceType", instances, (*AirTrafficControlServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficManagementServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficManagementServicePropertyType", instances, (*AirTrafficManagementServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficManagementServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficManagementServiceTimeSlicePropertyType", instances, (*AirTrafficManagementServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficManagementServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficManagementServiceTimeSliceType", instances, (*AirTrafficManagementServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirTrafficManagementServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirTrafficManagementServiceType", instances, (*AirTrafficManagementServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftCharacteristicPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftCharacteristicPropertyType", instances, (*AircraftCharacteristicPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftCharacteristicTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftCharacteristicType", instances, (*AircraftCharacteristicType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftGroundServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftGroundServicePropertyType", instances, (*AircraftGroundServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftGroundServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftGroundServiceTimeSlicePropertyType", instances, (*AircraftGroundServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftGroundServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftGroundServiceTimeSliceType", instances, (*AircraftGroundServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftGroundServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftGroundServiceType", instances, (*AircraftGroundServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandContaminationPropertyType", instances, (*AircraftStandContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandContaminationType", instances, (*AircraftStandContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandPropertyType", instances, (*AircraftStandPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandTimeSlicePropertyType", instances, (*AircraftStandTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandTimeSliceType", instances, (*AircraftStandTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftStandTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftStandType", instances, (*AircraftStandType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportClearanceServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportClearanceServicePropertyType", instances, (*AirportClearanceServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportClearanceServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportClearanceServiceTimeSlicePropertyType", instances, (*AirportClearanceServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportClearanceServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportClearanceServiceTimeSliceType", instances, (*AirportClearanceServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportClearanceServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportClearanceServiceType", instances, (*AirportClearanceServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportGroundServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportGroundServicePropertyType", instances, (*AirportGroundServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportAvailabilityPropertyType", instances, (*AirportHeliportAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportAvailabilityType", instances, (*AirportHeliportAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportCollocationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportCollocationPropertyType", instances, (*AirportHeliportCollocationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportCollocationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportCollocationTimeSlicePropertyType", instances, (*AirportHeliportCollocationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportCollocationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportCollocationTimeSliceType", instances, (*AirportHeliportCollocationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportCollocationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportCollocationType", instances, (*AirportHeliportCollocationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportContaminationPropertyType", instances, (*AirportHeliportContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportContaminationType", instances, (*AirportHeliportContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportPropertyType", instances, (*AirportHeliportPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportProtectionAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportProtectionAreaPropertyType", instances, (*AirportHeliportProtectionAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportResponsibilityOrganisationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportResponsibilityOrganisationPropertyType", instances, (*AirportHeliportResponsibilityOrganisationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportResponsibilityOrganisationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportResponsibilityOrganisationType", instances, (*AirportHeliportResponsibilityOrganisationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportTimeSlicePropertyType", instances, (*AirportHeliportTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportTimeSliceType", instances, (*AirportHeliportTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportType", instances, (*AirportHeliportType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportUsagePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportUsagePropertyType", instances, (*AirportHeliportUsagePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHeliportUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHeliportUsageType", instances, (*AirportHeliportUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHotSpotPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHotSpotPropertyType", instances, (*AirportHotSpotPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHotSpotTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHotSpotTimeSlicePropertyType", instances, (*AirportHotSpotTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHotSpotTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHotSpotTimeSliceType", instances, (*AirportHotSpotTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportHotSpotTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportHotSpotType", instances, (*AirportHotSpotType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportProtectionAreaMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportProtectionAreaMarkingPropertyType", instances, (*AirportProtectionAreaMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportProtectionAreaMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportProtectionAreaMarkingTimeSlicePropertyType", instances, (*AirportProtectionAreaMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportProtectionAreaMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportProtectionAreaMarkingTimeSliceType", instances, (*AirportProtectionAreaMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportProtectionAreaMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportProtectionAreaMarkingType", instances, (*AirportProtectionAreaMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportSuppliesServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServicePropertyType", instances, (*AirportSuppliesServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportSuppliesServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServiceTimeSlicePropertyType", instances, (*AirportSuppliesServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportSuppliesServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServiceTimeSliceType", instances, (*AirportSuppliesServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirportSuppliesServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirportSuppliesServiceType", instances, (*AirportSuppliesServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceActivationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceActivationPropertyType", instances, (*AirspaceActivationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceActivationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceActivationType", instances, (*AirspaceActivationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingPropertyType", instances, (*AirspaceBorderCrossingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingTimeSlicePropertyType", instances, (*AirspaceBorderCrossingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingTimeSliceType", instances, (*AirspaceBorderCrossingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceBorderCrossingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceBorderCrossingType", instances, (*AirspaceBorderCrossingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceGeometryComponentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceGeometryComponentPropertyType", instances, (*AirspaceGeometryComponentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceGeometryComponentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceGeometryComponentType", instances, (*AirspaceGeometryComponentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerClassPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerClassPropertyType", instances, (*AirspaceLayerClassPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerClassTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerClassType", instances, (*AirspaceLayerClassType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerPropertyType", instances, (*AirspaceLayerPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerType", instances, (*AirspaceLayerType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspacePropertyType", instances, (*AirspacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceTimeSlicePropertyType", instances, (*AirspaceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceTimeSliceType", instances, (*AirspaceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceType", instances, (*AirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumeDependencyPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumeDependencyPropertyType", instances, (*AirspaceVolumeDependencyPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumeDependencyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumeDependencyType", instances, (*AirspaceVolumeDependencyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumePropertyType", instances, (*AirspaceVolumePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumeType", instances, (*AirspaceVolumeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourcePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourcePropertyType", instances, (*AltimeterSourcePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceStatusPropertyType", instances, (*AltimeterSourceStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceStatusType", instances, (*AltimeterSourceStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceTimeSlicePropertyType", instances, (*AltimeterSourceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceTimeSliceType", instances, (*AltimeterSourceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltimeterSourceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltimeterSourceType", instances, (*AltimeterSourceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltitudeAdjustmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltitudeAdjustmentPropertyType", instances, (*AltitudeAdjustmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AltitudeAdjustmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AltitudeAdjustmentType", instances, (*AltitudeAdjustmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationPropertyType", instances, (*AngleIndicationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationTimeSlicePropertyType", instances, (*AngleIndicationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationTimeSliceType", instances, (*AngleIndicationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleIndicationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleIndicationType", instances, (*AngleIndicationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleUsePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleUsePropertyType", instances, (*AngleUsePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleUseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleUseType", instances, (*AngleUseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachAltitudeTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachAltitudeTablePropertyType", instances, (*ApproachAltitudeTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachAltitudeTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachAltitudeTableType", instances, (*ApproachAltitudeTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachConditionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachConditionPropertyType", instances, (*ApproachConditionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachConditionType", instances, (*ApproachConditionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachDistanceTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachDistanceTablePropertyType", instances, (*ApproachDistanceTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachDistanceTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachDistanceTableType", instances, (*ApproachDistanceTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLegPropertyType", instances, (*ApproachLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLightingSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLightingSystemPropertyType", instances, (*ApproachLightingSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLightingSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLightingSystemTimeSlicePropertyType", instances, (*ApproachLightingSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLightingSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLightingSystemTimeSliceType", instances, (*ApproachLightingSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachLightingSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachLightingSystemType", instances, (*ApproachLightingSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachTimingTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachTimingTablePropertyType", instances, (*ApproachTimingTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApproachTimingTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApproachTimingTableType", instances, (*ApproachTimingTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaAvailabilityPropertyType", instances, (*ApronAreaAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaAvailabilityType", instances, (*ApronAreaAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaUsagePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaUsagePropertyType", instances, (*ApronAreaUsagePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronAreaUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronAreaUsageType", instances, (*ApronAreaUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronContaminationPropertyType", instances, (*ApronContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronContaminationType", instances, (*ApronContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementPropertyType", instances, (*ApronElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementTimeSlicePropertyType", instances, (*ApronElementTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementTimeSliceType", instances, (*ApronElementTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronElementType", instances, (*ApronElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronLightSystemPropertyType", instances, (*ApronLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronLightSystemTimeSlicePropertyType", instances, (*ApronLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronLightSystemTimeSliceType", instances, (*ApronLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronLightSystemType", instances, (*ApronLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronMarkingPropertyType", instances, (*ApronMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronMarkingTimeSlicePropertyType", instances, (*ApronMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronMarkingTimeSliceType", instances, (*ApronMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronMarkingType", instances, (*ApronMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronPropertyType", instances, (*ApronPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronTimeSlicePropertyType", instances, (*ApronTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronTimeSliceType", instances, (*ApronTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ApronTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ApronType", instances, (*ApronType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrestingGearPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrestingGearPropertyType", instances, (*ArrestingGearPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrestingGearTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrestingGearTimeSlicePropertyType", instances, (*ArrestingGearTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrestingGearTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrestingGearTimeSliceType", instances, (*ArrestingGearTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrestingGearTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrestingGearType", instances, (*ArrestingGearType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalFeederLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalFeederLegPropertyType", instances, (*ArrivalFeederLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalFeederLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalFeederLegTimeSlicePropertyType", instances, (*ArrivalFeederLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalFeederLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalFeederLegTimeSliceType", instances, (*ArrivalFeederLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalFeederLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalFeederLegType", instances, (*ArrivalFeederLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalLegPropertyType", instances, (*ArrivalLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalLegTimeSlicePropertyType", instances, (*ArrivalLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalLegTimeSliceType", instances, (*ArrivalLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrivalLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrivalLegType", instances, (*ArrivalLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAerialRefuellingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAerialRefuellingPropertyType", instances, (*AuthorityForAerialRefuellingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAerialRefuellingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAerialRefuellingType", instances, (*AuthorityForAerialRefuellingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAirspacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAirspacePropertyType", instances, (*AuthorityForAirspacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAirspaceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAirspaceTimeSlicePropertyType", instances, (*AuthorityForAirspaceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAirspaceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAirspaceTimeSliceType", instances, (*AuthorityForAirspaceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAirspaceType", instances, (*AuthorityForAirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForNavaidEquipmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForNavaidEquipmentPropertyType", instances, (*AuthorityForNavaidEquipmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForNavaidEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForNavaidEquipmentType", instances, (*AuthorityForNavaidEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationStationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationStationPropertyType", instances, (*AuthorityForSpecialNavigationStationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationStationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationStationType", instances, (*AuthorityForSpecialNavigationStationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationSystemPropertyType", instances, (*AuthorityForSpecialNavigationSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForSpecialNavigationSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForSpecialNavigationSystemType", instances, (*AuthorityForSpecialNavigationSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AzimuthPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AzimuthPropertyType", instances, (*AzimuthPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AzimuthTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AzimuthTimeSlicePropertyType", instances, (*AzimuthTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AzimuthTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AzimuthTimeSliceType", instances, (*AzimuthTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AzimuthTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AzimuthType", instances, (*AzimuthType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CallsignDetailPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CallsignDetailPropertyType", instances, (*CallsignDetailPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CallsignDetailTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CallsignDetailType", instances, (*CallsignDetailType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChangeOverPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChangeOverPointPropertyType", instances, (*ChangeOverPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChangeOverPointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChangeOverPointTimeSlicePropertyType", instances, (*ChangeOverPointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChangeOverPointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChangeOverPointTimeSliceType", instances, (*ChangeOverPointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ChangeOverPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ChangeOverPointType", instances, (*ChangeOverPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointINSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointINSPropertyType", instances, (*CheckpointINSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointINSTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointINSTimeSlicePropertyType", instances, (*CheckpointINSTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointINSTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointINSTimeSliceType", instances, (*CheckpointINSTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointINSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointINSType", instances, (*CheckpointINSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointVORPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointVORPropertyType", instances, (*CheckpointVORPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointVORTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointVORTimeSlicePropertyType", instances, (*CheckpointVORTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointVORTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointVORTimeSliceType", instances, (*CheckpointVORTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CheckpointVORTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CheckpointVORType", instances, (*CheckpointVORType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CircleSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CircleSectorPropertyType", instances, (*CircleSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CircleSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CircleSectorType", instances, (*CircleSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingAreaPropertyType", instances, (*CirclingAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingAreaTimeSlicePropertyType", instances, (*CirclingAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingAreaTimeSliceType", instances, (*CirclingAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingAreaType", instances, (*CirclingAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingRestrictionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingRestrictionPropertyType", instances, (*CirclingRestrictionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CirclingRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CirclingRestrictionType", instances, (*CirclingRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CityPropertyType", instances, (*CityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CityType", instances, (*CityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConditionCombinationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConditionCombinationPropertyType", instances, (*ConditionCombinationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConditionCombinationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConditionCombinationType", instances, (*ConditionCombinationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ContactInformationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ContactInformationPropertyType", instances, (*ContactInformationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ContactInformationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ContactInformationType", instances, (*ContactInformationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurvePropertyType", instances, (*CurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurveType", instances, (*CurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMEPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMEPropertyType", instances, (*DMEPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMETimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMETimeSlicePropertyType", instances, (*DMETimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMETimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMETimeSliceType", instances, (*DMETimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMETypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMEType", instances, (*DMEType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaMarkingPropertyType", instances, (*DeicingAreaMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaMarkingTimeSlicePropertyType", instances, (*DeicingAreaMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaMarkingTimeSliceType", instances, (*DeicingAreaMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaMarkingType", instances, (*DeicingAreaMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaPropertyType", instances, (*DeicingAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaTimeSlicePropertyType", instances, (*DeicingAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaTimeSliceType", instances, (*DeicingAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DeicingAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DeicingAreaType", instances, (*DeicingAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureArrivalConditionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureArrivalConditionPropertyType", instances, (*DepartureArrivalConditionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureArrivalConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureArrivalConditionType", instances, (*DepartureArrivalConditionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureLegPropertyType", instances, (*DepartureLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureLegTimeSlicePropertyType", instances, (*DepartureLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureLegTimeSliceType", instances, (*DepartureLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DepartureLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DepartureLegType", instances, (*DepartureLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DesignatedPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointPropertyType", instances, (*DesignatedPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DesignatedPointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointTimeSlicePropertyType", instances, (*DesignatedPointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DesignatedPointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointTimeSliceType", instances, (*DesignatedPointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DesignatedPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DesignatedPointType", instances, (*DesignatedPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightClassPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightClassPropertyType", instances, (*DirectFlightClassPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightClassTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightClassType", instances, (*DirectFlightClassType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightPropertyType", instances, (*DirectFlightPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightSegmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightSegmentPropertyType", instances, (*DirectFlightSegmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectFlightSegmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectFlightSegmentType", instances, (*DirectFlightSegmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderPropertyType", instances, (*DirectionFinderPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderTimeSlicePropertyType", instances, (*DirectionFinderTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderTimeSliceType", instances, (*DirectionFinderTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionFinderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionFinderType", instances, (*DirectionFinderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationPropertyType", instances, (*DistanceIndicationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationTimeSlicePropertyType", instances, (*DistanceIndicationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationTimeSliceType", instances, (*DistanceIndicationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DistanceIndicationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DistanceIndicationType", instances, (*DistanceIndicationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedCurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedCurvePropertyType", instances, (*ElevatedCurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedCurveType", instances, (*ElevatedCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedPointPropertyType", instances, (*ElevatedPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedPointType", instances, (*ElevatedPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedSurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedSurfacePropertyType", instances, (*ElevatedSurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedSurfaceType", instances, (*ElevatedSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevationPropertyType", instances, (*ElevationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevationTimeSlicePropertyType", instances, (*ElevationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevationTimeSliceType", instances, (*ElevationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevationType", instances, (*ElevationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EnRouteSegmentPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EnRouteSegmentPointPropertyType", instances, (*EnRouteSegmentPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EnRouteSegmentPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EnRouteSegmentPointType", instances, (*EnRouteSegmentPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentColumnPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentColumnPropertyType", instances, (*EquipmentUnavailableAdjustmentColumnPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentColumnTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentColumnType", instances, (*EquipmentUnavailableAdjustmentColumnType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentPropertyType", instances, (*EquipmentUnavailableAdjustmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EquipmentUnavailableAdjustmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EquipmentUnavailableAdjustmentType", instances, (*EquipmentUnavailableAdjustmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FASDataBlockPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FASDataBlockPropertyType", instances, (*FASDataBlockPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FASDataBlockTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FASDataBlockType", instances, (*FASDataBlockType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalLegPropertyType", instances, (*FinalLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalLegTimeSlicePropertyType", instances, (*FinalLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalLegTimeSliceType", instances, (*FinalLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalLegType", instances, (*FinalLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalProfilePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalProfilePropertyType", instances, (*FinalProfilePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FinalProfileTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FinalProfileType", instances, (*FinalProfileType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FireFightingServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FireFightingServicePropertyType", instances, (*FireFightingServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FireFightingServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FireFightingServiceTimeSlicePropertyType", instances, (*FireFightingServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FireFightingServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FireFightingServiceTimeSliceType", instances, (*FireFightingServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FireFightingServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FireFightingServiceType", instances, (*FireFightingServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightCharacteristicPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightCharacteristicPropertyType", instances, (*FlightCharacteristicPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightCharacteristicTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightCharacteristicType", instances, (*FlightCharacteristicType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCircumstancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCircumstancePropertyType", instances, (*FlightConditionCircumstancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCircumstanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCircumstanceType", instances, (*FlightConditionCircumstanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCombinationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCombinationPropertyType", instances, (*FlightConditionCombinationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionCombinationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionCombinationType", instances, (*FlightConditionCombinationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionElementPropertyType", instances, (*FlightConditionElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightConditionElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightConditionElementType", instances, (*FlightConditionElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionLevelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionLevelPropertyType", instances, (*FlightRestrictionLevelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionLevelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionLevelType", instances, (*FlightRestrictionLevelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionPropertyType", instances, (*FlightRestrictionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionRoutePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionRoutePropertyType", instances, (*FlightRestrictionRoutePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionRouteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionRouteType", instances, (*FlightRestrictionRouteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionTimeSlicePropertyType", instances, (*FlightRestrictionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionTimeSliceType", instances, (*FlightRestrictionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRestrictionType", instances, (*FlightRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRoutingElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRoutingElementPropertyType", instances, (*FlightRoutingElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FlightRoutingElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FlightRoutingElementType", instances, (*FlightRoutingElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FloatingDockSitePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSitePropertyType", instances, (*FloatingDockSitePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FloatingDockSiteTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSiteTimeSlicePropertyType", instances, (*FloatingDockSiteTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FloatingDockSiteTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSiteTimeSliceType", instances, (*FloatingDockSiteTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FloatingDockSiteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FloatingDockSiteType", instances, (*FloatingDockSiteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FuelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FuelPropertyType", instances, (*FuelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FuelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FuelType", instances, (*FuelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeoBorderPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeoBorderPropertyType", instances, (*GeoBorderPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeoBorderTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeoBorderTimeSlicePropertyType", instances, (*GeoBorderTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeoBorderTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeoBorderTimeSliceType", instances, (*GeoBorderTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeoBorderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeoBorderType", instances, (*GeoBorderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GlidepathPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GlidepathPropertyType", instances, (*GlidepathPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GlidepathTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GlidepathTimeSlicePropertyType", instances, (*GlidepathTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GlidepathTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GlidepathTimeSliceType", instances, (*GlidepathTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GlidepathTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GlidepathType", instances, (*GlidepathType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundLightSystemPropertyType", instances, (*GroundLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundLightingAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundLightingAvailabilityPropertyType", instances, (*GroundLightingAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundLightingAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundLightingAvailabilityType", instances, (*GroundLightingAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundTrafficControlServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundTrafficControlServicePropertyType", instances, (*GroundTrafficControlServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundTrafficControlServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundTrafficControlServiceTimeSlicePropertyType", instances, (*GroundTrafficControlServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundTrafficControlServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundTrafficControlServiceTimeSliceType", instances, (*GroundTrafficControlServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GroundTrafficControlServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GroundTrafficControlServiceType", instances, (*GroundTrafficControlServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineLightSystemPropertyType", instances, (*GuidanceLineLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineLightSystemTimeSlicePropertyType", instances, (*GuidanceLineLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineLightSystemTimeSliceType", instances, (*GuidanceLineLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineLightSystemType", instances, (*GuidanceLineLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineMarkingPropertyType", instances, (*GuidanceLineMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineMarkingTimeSlicePropertyType", instances, (*GuidanceLineMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineMarkingTimeSliceType", instances, (*GuidanceLineMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineMarkingType", instances, (*GuidanceLineMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLinePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLinePropertyType", instances, (*GuidanceLinePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineTimeSlicePropertyType", instances, (*GuidanceLineTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineTimeSliceType", instances, (*GuidanceLineTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GuidanceLineTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GuidanceLineType", instances, (*GuidanceLineType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingAssessmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingAssessmentPropertyType", instances, (*HoldingAssessmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingAssessmentTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingAssessmentTimeSlicePropertyType", instances, (*HoldingAssessmentTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingAssessmentTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingAssessmentTimeSliceType", instances, (*HoldingAssessmentTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingAssessmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingAssessmentType", instances, (*HoldingAssessmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDistancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDistancePropertyType", instances, (*HoldingPatternDistancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDistanceType", instances, (*HoldingPatternDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDurationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDurationPropertyType", instances, (*HoldingPatternDurationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternDurationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternDurationType", instances, (*HoldingPatternDurationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternPropertyType", instances, (*HoldingPatternPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternTimeSlicePropertyType", instances, (*HoldingPatternTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternTimeSliceType", instances, (*HoldingPatternTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingPatternTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingPatternType", instances, (*HoldingPatternType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingUsePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingUsePropertyType", instances, (*HoldingUsePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoldingUseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoldingUseType", instances, (*HoldingUseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InformationServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServicePropertyType", instances, (*InformationServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InformationServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServiceTimeSlicePropertyType", instances, (*InformationServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InformationServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServiceTimeSliceType", instances, (*InformationServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InformationServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InformationServiceType", instances, (*InformationServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialLegPropertyType", instances, (*InitialLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialLegTimeSlicePropertyType", instances, (*InitialLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialLegTimeSliceType", instances, (*InitialLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InitialLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InitialLegType", instances, (*InitialLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedurePropertyType", instances, (*InstrumentApproachProcedurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedureTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedureTimeSlicePropertyType", instances, (*InstrumentApproachProcedureTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedureTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedureTimeSliceType", instances, (*InstrumentApproachProcedureTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentApproachProcedureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentApproachProcedureType", instances, (*InstrumentApproachProcedureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IntermediateLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IntermediateLegPropertyType", instances, (*IntermediateLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IntermediateLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IntermediateLegTimeSlicePropertyType", instances, (*IntermediateLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IntermediateLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IntermediateLegTimeSliceType", instances, (*IntermediateLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IntermediateLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IntermediateLegType", instances, (*IntermediateLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LandingTakeoffAreaCollectionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LandingTakeoffAreaCollectionPropertyType", instances, (*LandingTakeoffAreaCollectionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LandingTakeoffAreaCollectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LandingTakeoffAreaCollectionType", instances, (*LandingTakeoffAreaCollectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightActivationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightActivationPropertyType", instances, (*LightActivationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightActivationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightActivationType", instances, (*LightActivationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementPropertyType", instances, (*LightElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementStatusPropertyType", instances, (*LightElementStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementStatusType", instances, (*LightElementStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LightElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LightElementType", instances, (*LightElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinguisticNotePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinguisticNotePropertyType", instances, (*LinguisticNotePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinguisticNoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinguisticNoteType", instances, (*LinguisticNoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalizerPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalizerPropertyType", instances, (*LocalizerPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalizerTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalizerTimeSlicePropertyType", instances, (*LocalizerTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalizerTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalizerTimeSliceType", instances, (*LocalizerTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalizerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalizerType", instances, (*LocalizerType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ManoeuvringAreaAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaAvailabilityPropertyType", instances, (*ManoeuvringAreaAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ManoeuvringAreaAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaAvailabilityType", instances, (*ManoeuvringAreaAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ManoeuvringAreaUsagePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaUsagePropertyType", instances, (*ManoeuvringAreaUsagePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ManoeuvringAreaUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ManoeuvringAreaUsageType", instances, (*ManoeuvringAreaUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkerBeaconPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkerBeaconPropertyType", instances, (*MarkerBeaconPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkerBeaconTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkerBeaconTimeSlicePropertyType", instances, (*MarkerBeaconTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkerBeaconTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkerBeaconTimeSliceType", instances, (*MarkerBeaconTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkerBeaconTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkerBeaconType", instances, (*MarkerBeaconType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingBuoyPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingBuoyPropertyType", instances, (*MarkingBuoyPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingBuoyTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingBuoyTimeSlicePropertyType", instances, (*MarkingBuoyTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingBuoyTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingBuoyTimeSliceType", instances, (*MarkingBuoyTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingBuoyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingBuoyType", instances, (*MarkingBuoyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingElementPropertyType", instances, (*MarkingElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingElementType", instances, (*MarkingElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MarkingPropertyType", instances, (*MarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MeteorologyPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MeteorologyPropertyType", instances, (*MeteorologyPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MeteorologyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MeteorologyType", instances, (*MeteorologyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinimaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MinimaPropertyType", instances, (*MinimaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinimaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MinimaType", instances, (*MinimaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachGroupPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachGroupPropertyType", instances, (*MissedApproachGroupPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachGroupTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachGroupType", instances, (*MissedApproachGroupType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachLegPropertyType", instances, (*MissedApproachLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachLegTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachLegTimeSlicePropertyType", instances, (*MissedApproachLegTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachLegTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachLegTimeSliceType", instances, (*MissedApproachLegTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MissedApproachLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MissedApproachLegType", instances, (*MissedApproachLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NDBPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NDBPropertyType", instances, (*NDBPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NDBTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NDBTimeSlicePropertyType", instances, (*NDBTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NDBTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NDBTimeSliceType", instances, (*NDBTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NDBTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NDBType", instances, (*NDBType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidComponentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidComponentPropertyType", instances, (*NavaidComponentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidComponentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidComponentType", instances, (*NavaidComponentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentDistancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentDistancePropertyType", instances, (*NavaidEquipmentDistancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentDistanceType", instances, (*NavaidEquipmentDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentMonitoringPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentMonitoringPropertyType", instances, (*NavaidEquipmentMonitoringPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentMonitoringTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentMonitoringType", instances, (*NavaidEquipmentMonitoringType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidEquipmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidEquipmentPropertyType", instances, (*NavaidEquipmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidOperationalStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidOperationalStatusPropertyType", instances, (*NavaidOperationalStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidOperationalStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidOperationalStatusType", instances, (*NavaidOperationalStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidPropertyType", instances, (*NavaidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidTimeSlicePropertyType", instances, (*NavaidTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidTimeSliceType", instances, (*NavaidTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavaidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavaidType", instances, (*NavaidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaPropertyType", instances, (*NavigationAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaRestrictionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaRestrictionPropertyType", instances, (*NavigationAreaRestrictionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaRestrictionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaRestrictionTimeSlicePropertyType", instances, (*NavigationAreaRestrictionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaRestrictionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaRestrictionTimeSliceType", instances, (*NavigationAreaRestrictionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaRestrictionType", instances, (*NavigationAreaRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaSectorPropertyType", instances, (*NavigationAreaSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaSectorType", instances, (*NavigationAreaSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaTimeSlicePropertyType", instances, (*NavigationAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaTimeSliceType", instances, (*NavigationAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationAreaType", instances, (*NavigationAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NavigationSystemCheckpointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NavigationSystemCheckpointPropertyType", instances, (*NavigationSystemCheckpointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NitrogenPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NitrogenPropertyType", instances, (*NitrogenPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NitrogenTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NitrogenType", instances, (*NitrogenType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NonMovementAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonMovementAreaPropertyType", instances, (*NonMovementAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NonMovementAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonMovementAreaTimeSlicePropertyType", instances, (*NonMovementAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NonMovementAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonMovementAreaTimeSliceType", instances, (*NonMovementAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NonMovementAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonMovementAreaType", instances, (*NonMovementAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NotePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NotePropertyType", instances, (*NotePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoteType", instances, (*NoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAreaPropertyType", instances, (*ObstacleAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAreaTimeSlicePropertyType", instances, (*ObstacleAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAreaTimeSliceType", instances, (*ObstacleAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAreaType", instances, (*ObstacleAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAssessmentAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAssessmentAreaPropertyType", instances, (*ObstacleAssessmentAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstacleAssessmentAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstacleAssessmentAreaType", instances, (*ObstacleAssessmentAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstaclePlacementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstaclePlacementPropertyType", instances, (*ObstaclePlacementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstaclePlacementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstaclePlacementType", instances, (*ObstaclePlacementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstructionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstructionPropertyType", instances, (*ObstructionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObstructionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObstructionType", instances, (*ObstructionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OilPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OilPropertyType", instances, (*OilPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OilTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OilType", instances, (*OilType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OnlineContactPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OnlineContactPropertyType", instances, (*OnlineContactPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OnlineContactTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OnlineContactType", instances, (*OnlineContactType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityAssociationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityAssociationPropertyType", instances, (*OrganisationAuthorityAssociationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityAssociationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityAssociationType", instances, (*OrganisationAuthorityAssociationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityPropertyType", instances, (*OrganisationAuthorityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityTimeSlicePropertyType", instances, (*OrganisationAuthorityTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityTimeSliceType", instances, (*OrganisationAuthorityTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrganisationAuthorityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityType", instances, (*OrganisationAuthorityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OxygenPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OxygenPropertyType", instances, (*OxygenPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OxygenTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OxygenType", instances, (*OxygenType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerLoadingBridgePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerLoadingBridgePropertyType", instances, (*PassengerLoadingBridgePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerLoadingBridgeTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerLoadingBridgeTimeSlicePropertyType", instances, (*PassengerLoadingBridgeTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerLoadingBridgeTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerLoadingBridgeTimeSliceType", instances, (*PassengerLoadingBridgeTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerLoadingBridgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerLoadingBridgeType", instances, (*PassengerLoadingBridgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerServicePropertyType", instances, (*PassengerServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerServiceTimeSlicePropertyType", instances, (*PassengerServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerServiceTimeSliceType", instances, (*PassengerServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassengerServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassengerServiceType", instances, (*PassengerServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PilotControlledLightingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PilotControlledLightingPropertyType", instances, (*PilotControlledLightingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PilotControlledLightingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PilotControlledLightingTimeSlicePropertyType", instances, (*PilotControlledLightingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PilotControlledLightingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PilotControlledLightingTimeSliceType", instances, (*PilotControlledLightingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PilotControlledLightingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PilotControlledLightingType", instances, (*PilotControlledLightingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointPropertyType", instances, (*PointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointReferencePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointReferencePropertyType", instances, (*PointReferencePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointReferenceType", instances, (*PointReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointType", instances, (*PointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PostalAddressPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PostalAddressPropertyType", instances, (*PostalAddressPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PostalAddressTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PostalAddressType", instances, (*PostalAddressType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrecisionApproachRadarPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrecisionApproachRadarPropertyType", instances, (*PrecisionApproachRadarPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrecisionApproachRadarTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrecisionApproachRadarTimeSlicePropertyType", instances, (*PrecisionApproachRadarTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrecisionApproachRadarTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrecisionApproachRadarTimeSliceType", instances, (*PrecisionApproachRadarTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrecisionApproachRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrecisionApproachRadarType", instances, (*PrecisionApproachRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimarySurveillanceRadarPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimarySurveillanceRadarPropertyType", instances, (*PrimarySurveillanceRadarPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimarySurveillanceRadarTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimarySurveillanceRadarTimeSlicePropertyType", instances, (*PrimarySurveillanceRadarTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimarySurveillanceRadarTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimarySurveillanceRadarTimeSliceType", instances, (*PrimarySurveillanceRadarTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimarySurveillanceRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimarySurveillanceRadarType", instances, (*PrimarySurveillanceRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureAvailabilityPropertyType", instances, (*ProcedureAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureAvailabilityType", instances, (*ProcedureAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureDMEPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureDMEPropertyType", instances, (*ProcedureDMEPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureDMETimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureDMETimeSlicePropertyType", instances, (*ProcedureDMETimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureDMETimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureDMETimeSliceType", instances, (*ProcedureDMETimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureDMETypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureDMEType", instances, (*ProcedureDMEType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedurePropertyType", instances, (*ProcedurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionLegPropertyType", instances, (*ProcedureTransitionLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionLegTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionLegType", instances, (*ProcedureTransitionLegType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionPropertyType", instances, (*ProcedureTransitionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedureTransitionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedureTransitionType", instances, (*ProcedureTransitionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PropertiesWithSchedulePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PropertiesWithSchedulePropertyType", instances, (*PropertiesWithSchedulePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarComponentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarComponentPropertyType", instances, (*RadarComponentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarComponentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarComponentType", instances, (*RadarComponentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarEquipmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarEquipmentPropertyType", instances, (*RadarEquipmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemPropertyType", instances, (*RadarSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemTimeSlicePropertyType", instances, (*RadarSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemTimeSliceType", instances, (*RadarSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadarSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadarSystemType", instances, (*RadarSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelPropertyType", instances, (*RadioCommunicationChannelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelTimeSlicePropertyType", instances, (*RadioCommunicationChannelTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelTimeSliceType", instances, (*RadioCommunicationChannelTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationChannelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationChannelType", instances, (*RadioCommunicationChannelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationOperationalStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationOperationalStatusPropertyType", instances, (*RadioCommunicationOperationalStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioCommunicationOperationalStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioCommunicationOperationalStatusType", instances, (*RadioCommunicationOperationalStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioFrequencyAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioFrequencyAreaPropertyType", instances, (*RadioFrequencyAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioFrequencyAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioFrequencyAreaTimeSlicePropertyType", instances, (*RadioFrequencyAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioFrequencyAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioFrequencyAreaTimeSliceType", instances, (*RadioFrequencyAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RadioFrequencyAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RadioFrequencyAreaType", instances, (*RadioFrequencyAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ReflectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ReflectorPropertyType", instances, (*ReflectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ReflectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ReflectorType", instances, (*ReflectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RidgePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RidgePropertyType", instances, (*RidgePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RidgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RidgeType", instances, (*RidgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoadPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoadPropertyType", instances, (*RoadPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoadTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoadTimeSlicePropertyType", instances, (*RoadTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoadTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoadTimeSliceType", instances, (*RoadTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoadTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoadType", instances, (*RoadType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteAvailabilityPropertyType", instances, (*RouteAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteAvailabilityType", instances, (*RouteAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteDMEPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteDMEPropertyType", instances, (*RouteDMEPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteDMETimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteDMETimeSlicePropertyType", instances, (*RouteDMETimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteDMETimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteDMETimeSliceType", instances, (*RouteDMETimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteDMETypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteDMEType", instances, (*RouteDMEType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoutePortionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoutePortionPropertyType", instances, (*RoutePortionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoutePortionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoutePortionType", instances, (*RoutePortionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoutePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoutePropertyType", instances, (*RoutePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteSegmentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteSegmentPropertyType", instances, (*RouteSegmentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteSegmentTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteSegmentTimeSlicePropertyType", instances, (*RouteSegmentTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteSegmentTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteSegmentTimeSliceType", instances, (*RouteSegmentTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteSegmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteSegmentType", instances, (*RouteSegmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteTimeSlicePropertyType", instances, (*RouteTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteTimeSliceType", instances, (*RouteTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RouteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RouteType", instances, (*RouteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RulesProceduresPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RulesProceduresPropertyType", instances, (*RulesProceduresPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RulesProceduresTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RulesProceduresTimeSlicePropertyType", instances, (*RulesProceduresTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RulesProceduresTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RulesProceduresTimeSliceType", instances, (*RulesProceduresTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RulesProceduresTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RulesProceduresType", instances, (*RulesProceduresType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayBlastPadPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayBlastPadPropertyType", instances, (*RunwayBlastPadPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayBlastPadTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayBlastPadTimeSlicePropertyType", instances, (*RunwayBlastPadTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayBlastPadTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayBlastPadTimeSliceType", instances, (*RunwayBlastPadTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayBlastPadTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayBlastPadType", instances, (*RunwayBlastPadType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayCentrelinePointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointPropertyType", instances, (*RunwayCentrelinePointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayCentrelinePointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointTimeSlicePropertyType", instances, (*RunwayCentrelinePointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayCentrelinePointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointTimeSliceType", instances, (*RunwayCentrelinePointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayCentrelinePointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayCentrelinePointType", instances, (*RunwayCentrelinePointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayContaminationPropertyType", instances, (*RunwayContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayContaminationType", instances, (*RunwayContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistancePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistancePropertyType", instances, (*RunwayDeclaredDistancePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistanceType", instances, (*RunwayDeclaredDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistanceValuePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistanceValuePropertyType", instances, (*RunwayDeclaredDistanceValuePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDeclaredDistanceValueTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDeclaredDistanceValueType", instances, (*RunwayDeclaredDistanceValueType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionLightSystemPropertyType", instances, (*RunwayDirectionLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionLightSystemTimeSlicePropertyType", instances, (*RunwayDirectionLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionLightSystemTimeSliceType", instances, (*RunwayDirectionLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionLightSystemType", instances, (*RunwayDirectionLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionPropertyType", instances, (*RunwayDirectionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionTimeSlicePropertyType", instances, (*RunwayDirectionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionTimeSliceType", instances, (*RunwayDirectionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayDirectionType", instances, (*RunwayDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementPropertyType", instances, (*RunwayElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementTimeSlicePropertyType", instances, (*RunwayElementTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementTimeSliceType", instances, (*RunwayElementTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayElementType", instances, (*RunwayElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayMarkingPropertyType", instances, (*RunwayMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayMarkingTimeSlicePropertyType", instances, (*RunwayMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayMarkingTimeSliceType", instances, (*RunwayMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayMarkingType", instances, (*RunwayMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayPropertyType", instances, (*RunwayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaLightSystemPropertyType", instances, (*RunwayProtectAreaLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaLightSystemTimeSlicePropertyType", instances, (*RunwayProtectAreaLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaLightSystemTimeSliceType", instances, (*RunwayProtectAreaLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaLightSystemType", instances, (*RunwayProtectAreaLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaPropertyType", instances, (*RunwayProtectAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaTimeSlicePropertyType", instances, (*RunwayProtectAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaTimeSliceType", instances, (*RunwayProtectAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayProtectAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayProtectAreaType", instances, (*RunwayProtectAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwaySectionContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwaySectionContaminationPropertyType", instances, (*RunwaySectionContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwaySectionContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwaySectionContaminationType", instances, (*RunwaySectionContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayTimeSlicePropertyType", instances, (*RunwayTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayTimeSliceType", instances, (*RunwayTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayType", instances, (*RunwayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayVisualRangePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayVisualRangePropertyType", instances, (*RunwayVisualRangePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayVisualRangeTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayVisualRangeTimeSlicePropertyType", instances, (*RunwayVisualRangeTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayVisualRangeTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayVisualRangeTimeSliceType", instances, (*RunwayVisualRangeTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RunwayVisualRangeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RunwayVisualRangeType", instances, (*RunwayVisualRangeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SDFPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SDFPropertyType", instances, (*SDFPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SDFTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SDFTimeSlicePropertyType", instances, (*SDFTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SDFTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SDFTimeSliceType", instances, (*SDFTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SDFTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SDFType", instances, (*SDFType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaPropertyType", instances, (*SafeAltitudeAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaSectorPropertyType", instances, (*SafeAltitudeAreaSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaSectorType", instances, (*SafeAltitudeAreaSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaTimeSlicePropertyType", instances, (*SafeAltitudeAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaTimeSliceType", instances, (*SafeAltitudeAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SafeAltitudeAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SafeAltitudeAreaType", instances, (*SafeAltitudeAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaPropertyType", instances, (*SeaplaneLandingAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaTimeSlicePropertyType", instances, (*SeaplaneLandingAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaTimeSliceType", instances, (*SeaplaneLandingAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneLandingAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneLandingAreaType", instances, (*SeaplaneLandingAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSitePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSitePropertyType", instances, (*SeaplaneRampSitePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSiteTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSiteTimeSlicePropertyType", instances, (*SeaplaneRampSiteTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSiteTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSiteTimeSliceType", instances, (*SeaplaneRampSiteTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SeaplaneRampSiteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SeaplaneRampSiteType", instances, (*SeaplaneRampSiteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SearchRescueServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SearchRescueServicePropertyType", instances, (*SearchRescueServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SearchRescueServiceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SearchRescueServiceTimeSlicePropertyType", instances, (*SearchRescueServiceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SearchRescueServiceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SearchRescueServiceTimeSliceType", instances, (*SearchRescueServiceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SearchRescueServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SearchRescueServiceType", instances, (*SearchRescueServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondarySurveillanceRadarPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondarySurveillanceRadarPropertyType", instances, (*SecondarySurveillanceRadarPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondarySurveillanceRadarTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondarySurveillanceRadarTimeSlicePropertyType", instances, (*SecondarySurveillanceRadarTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondarySurveillanceRadarTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondarySurveillanceRadarTimeSliceType", instances, (*SecondarySurveillanceRadarTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondarySurveillanceRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondarySurveillanceRadarType", instances, (*SecondarySurveillanceRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SectorDesignPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SectorDesignPropertyType", instances, (*SectorDesignPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SectorDesignTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SectorDesignType", instances, (*SectorDesignType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SegmentLegPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SegmentLegPropertyType", instances, (*SegmentLegPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SegmentPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SegmentPointPropertyType", instances, (*SegmentPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ServiceOperationalStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ServiceOperationalStatusPropertyType", instances, (*ServiceOperationalStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ServiceOperationalStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ServiceOperationalStatusType", instances, (*ServiceOperationalStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ServicePropertyType", instances, (*ServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SignificantPointInAirspacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SignificantPointInAirspacePropertyType", instances, (*SignificantPointInAirspacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SignificantPointInAirspaceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SignificantPointInAirspaceTimeSlicePropertyType", instances, (*SignificantPointInAirspaceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SignificantPointInAirspaceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SignificantPointInAirspaceTimeSliceType", instances, (*SignificantPointInAirspaceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SignificantPointInAirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SignificantPointInAirspaceType", instances, (*SignificantPointInAirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialDatePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialDatePropertyType", instances, (*SpecialDatePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialDateTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialDateTimeSlicePropertyType", instances, (*SpecialDateTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialDateTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialDateTimeSliceType", instances, (*SpecialDateTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialDateTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialDateType", instances, (*SpecialDateType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationPropertyType", instances, (*SpecialNavigationStationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationStatusPropertyType", instances, (*SpecialNavigationStationStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationStatusType", instances, (*SpecialNavigationStationStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationTimeSlicePropertyType", instances, (*SpecialNavigationStationTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationTimeSliceType", instances, (*SpecialNavigationStationTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationStationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationStationType", instances, (*SpecialNavigationStationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemPropertyType", instances, (*SpecialNavigationSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemTimeSlicePropertyType", instances, (*SpecialNavigationSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemTimeSliceType", instances, (*SpecialNavigationSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SpecialNavigationSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SpecialNavigationSystemType", instances, (*SpecialNavigationSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandMarkingPropertyType", instances, (*StandMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandMarkingTimeSlicePropertyType", instances, (*StandMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandMarkingTimeSliceType", instances, (*StandMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandMarkingType", instances, (*StandMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentArrivalPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalPropertyType", instances, (*StandardInstrumentArrivalPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentArrivalTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalTimeSlicePropertyType", instances, (*StandardInstrumentArrivalTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentArrivalTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalTimeSliceType", instances, (*StandardInstrumentArrivalTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentArrivalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentArrivalType", instances, (*StandardInstrumentArrivalType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentDeparturePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDeparturePropertyType", instances, (*StandardInstrumentDeparturePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentDepartureTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDepartureTimeSlicePropertyType", instances, (*StandardInstrumentDepartureTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentDepartureTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDepartureTimeSliceType", instances, (*StandardInstrumentDepartureTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardInstrumentDepartureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardInstrumentDepartureType", instances, (*StandardInstrumentDepartureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnPropertyType", instances, (*StandardLevelColumnPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnTimeSlicePropertyType", instances, (*StandardLevelColumnTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnTimeSliceType", instances, (*StandardLevelColumnTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnType", instances, (*StandardLevelColumnType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelPropertyType", instances, (*StandardLevelPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelSectorPropertyType", instances, (*StandardLevelSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelSectorTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelSectorTimeSlicePropertyType", instances, (*StandardLevelSectorTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelSectorTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelSectorTimeSliceType", instances, (*StandardLevelSectorTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelSectorType", instances, (*StandardLevelSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTablePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTablePropertyType", instances, (*StandardLevelTablePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTableTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTableTimeSlicePropertyType", instances, (*StandardLevelTableTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTableTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTableTimeSliceType", instances, (*StandardLevelTableTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelTableType", instances, (*StandardLevelTableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelType", instances, (*StandardLevelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceCharacteristicsPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceCharacteristicsPropertyType", instances, (*SurfaceCharacteristicsPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceCharacteristicsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceCharacteristicsType", instances, (*SurfaceCharacteristicsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceContaminationLayerPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceContaminationLayerPropertyType", instances, (*SurfaceContaminationLayerPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceContaminationLayerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceContaminationLayerType", instances, (*SurfaceContaminationLayerType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceContaminationPropertyType", instances, (*SurfaceContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfacePropertyType", instances, (*SurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceType", instances, (*SurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveillanceGroundStationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveillanceGroundStationPropertyType", instances, (*SurveillanceGroundStationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveillanceGroundStationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveillanceGroundStationType", instances, (*SurveillanceGroundStationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveillanceRadarPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveillanceRadarPropertyType", instances, (*SurveillanceRadarPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveyControlPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveyControlPointPropertyType", instances, (*SurveyControlPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveyControlPointTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveyControlPointTimeSlicePropertyType", instances, (*SurveyControlPointTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveyControlPointTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveyControlPointTimeSliceType", instances, (*SurveyControlPointTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurveyControlPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurveyControlPointType", instances, (*SurveyControlPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TACANPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TACANPropertyType", instances, (*TACANPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TACANTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TACANTimeSlicePropertyType", instances, (*TACANTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TACANTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TACANTimeSliceType", instances, (*TACANTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TACANTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TACANType", instances, (*TACANType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionLightSystemPropertyType", instances, (*TaxiHoldingPositionLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionLightSystemTimeSlicePropertyType", instances, (*TaxiHoldingPositionLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionLightSystemTimeSliceType", instances, (*TaxiHoldingPositionLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionLightSystemType", instances, (*TaxiHoldingPositionLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionMarkingPropertyType", instances, (*TaxiHoldingPositionMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionMarkingTimeSlicePropertyType", instances, (*TaxiHoldingPositionMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionMarkingTimeSliceType", instances, (*TaxiHoldingPositionMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionMarkingType", instances, (*TaxiHoldingPositionMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionPropertyType", instances, (*TaxiHoldingPositionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionTimeSlicePropertyType", instances, (*TaxiHoldingPositionTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionTimeSliceType", instances, (*TaxiHoldingPositionTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiHoldingPositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiHoldingPositionType", instances, (*TaxiHoldingPositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayContaminationPropertyType", instances, (*TaxiwayContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayContaminationType", instances, (*TaxiwayContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementPropertyType", instances, (*TaxiwayElementPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementTimeSlicePropertyType", instances, (*TaxiwayElementTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementTimeSliceType", instances, (*TaxiwayElementTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayElementType", instances, (*TaxiwayElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayLightSystemPropertyType", instances, (*TaxiwayLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayLightSystemTimeSlicePropertyType", instances, (*TaxiwayLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayLightSystemTimeSliceType", instances, (*TaxiwayLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayLightSystemType", instances, (*TaxiwayLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayMarkingPropertyType", instances, (*TaxiwayMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayMarkingTimeSlicePropertyType", instances, (*TaxiwayMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayMarkingTimeSliceType", instances, (*TaxiwayMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayMarkingType", instances, (*TaxiwayMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayPropertyType", instances, (*TaxiwayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayTimeSlicePropertyType", instances, (*TaxiwayTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayTimeSliceType", instances, (*TaxiwayTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TaxiwayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TaxiwayType", instances, (*TaxiwayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TelephoneContactPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TelephoneContactPropertyType", instances, (*TelephoneContactPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TelephoneContactTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TelephoneContactType", instances, (*TelephoneContactType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaPropertyType", instances, (*TerminalArrivalAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaSectorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaSectorPropertyType", instances, (*TerminalArrivalAreaSectorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaSectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaSectorType", instances, (*TerminalArrivalAreaSectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaTimeSlicePropertyType", instances, (*TerminalArrivalAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaTimeSliceType", instances, (*TerminalArrivalAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalArrivalAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalArrivalAreaType", instances, (*TerminalArrivalAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalSegmentPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalSegmentPointPropertyType", instances, (*TerminalSegmentPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TerminalSegmentPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TerminalSegmentPointType", instances, (*TerminalSegmentPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimesheetPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimesheetPropertyType", instances, (*TimesheetPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimesheetTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimesheetType", instances, (*TimesheetType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffContaminationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffContaminationPropertyType", instances, (*TouchDownLiftOffContaminationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffContaminationType", instances, (*TouchDownLiftOffContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffLightSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffLightSystemPropertyType", instances, (*TouchDownLiftOffLightSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffLightSystemTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffLightSystemTimeSlicePropertyType", instances, (*TouchDownLiftOffLightSystemTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffLightSystemTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffLightSystemTimeSliceType", instances, (*TouchDownLiftOffLightSystemTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffLightSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffLightSystemType", instances, (*TouchDownLiftOffLightSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffMarkingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffMarkingPropertyType", instances, (*TouchDownLiftOffMarkingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffMarkingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffMarkingTimeSlicePropertyType", instances, (*TouchDownLiftOffMarkingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffMarkingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffMarkingTimeSliceType", instances, (*TouchDownLiftOffMarkingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffMarkingType", instances, (*TouchDownLiftOffMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffPropertyType", instances, (*TouchDownLiftOffPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffSafeAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffSafeAreaPropertyType", instances, (*TouchDownLiftOffSafeAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffSafeAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffSafeAreaTimeSlicePropertyType", instances, (*TouchDownLiftOffSafeAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffSafeAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffSafeAreaTimeSliceType", instances, (*TouchDownLiftOffSafeAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffSafeAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffSafeAreaType", instances, (*TouchDownLiftOffSafeAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffTimeSlicePropertyType", instances, (*TouchDownLiftOffTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffTimeSliceType", instances, (*TouchDownLiftOffTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TouchDownLiftOffTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TouchDownLiftOffType", instances, (*TouchDownLiftOffType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TrafficSeparationServicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TrafficSeparationServicePropertyType", instances, (*TrafficSeparationServicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitAvailabilityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitAvailabilityPropertyType", instances, (*UnitAvailabilityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitAvailabilityType", instances, (*UnitAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitDependencyPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitDependencyPropertyType", instances, (*UnitDependencyPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitDependencyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitDependencyType", instances, (*UnitDependencyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitPropertyType", instances, (*UnitPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitTimeSlicePropertyType", instances, (*UnitTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitTimeSliceType", instances, (*UnitTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitType", instances, (*UnitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnplannedHoldingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingPropertyType", instances, (*UnplannedHoldingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnplannedHoldingTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingTimeSlicePropertyType", instances, (*UnplannedHoldingTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnplannedHoldingTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingTimeSliceType", instances, (*UnplannedHoldingTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnplannedHoldingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnplannedHoldingType", instances, (*UnplannedHoldingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UsageConditionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UsageConditionPropertyType", instances, (*UsageConditionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORPropertyType", instances, (*VORPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORTimeSlicePropertyType", instances, (*VORTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORTimeSliceType", instances, (*VORTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VORTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VORType", instances, (*VORType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructureLightingStatusPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructureLightingStatusPropertyType", instances, (*VerticalStructureLightingStatusPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructureLightingStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructureLightingStatusType", instances, (*VerticalStructureLightingStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructurePartPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructurePartPropertyType", instances, (*VerticalStructurePartPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructurePartTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructurePartType", instances, (*VerticalStructurePartType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructurePropertyType", instances, (*VerticalStructurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructureTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructureTimeSlicePropertyType", instances, (*VerticalStructureTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructureTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructureTimeSliceType", instances, (*VerticalStructureTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalStructureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalStructureType", instances, (*VerticalStructureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VisualGlideSlopeIndicatorPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VisualGlideSlopeIndicatorPropertyType", instances, (*VisualGlideSlopeIndicatorPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VisualGlideSlopeIndicatorTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VisualGlideSlopeIndicatorTimeSlicePropertyType", instances, (*VisualGlideSlopeIndicatorTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VisualGlideSlopeIndicatorTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VisualGlideSlopeIndicatorTimeSliceType", instances, (*VisualGlideSlopeIndicatorTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VisualGlideSlopeIndicatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VisualGlideSlopeIndicatorType", instances, (*VisualGlideSlopeIndicatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkAreaPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkAreaPropertyType", instances, (*WorkAreaPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkAreaTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkAreaTimeSlicePropertyType", instances, (*WorkAreaTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkAreaTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkAreaTimeSliceType", instances, (*WorkAreaTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkAreaType", instances, (*WorkAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkareaActivityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkareaActivityPropertyType", instances, (*WorkareaActivityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WorkareaActivityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WorkareaActivityType", instances, (*WorkareaActivityType)(nil).GongGetFieldHeaders(), addIDs)
		}
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab GongExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// GongTabulator is an interface for writing to a table strings
type GongTabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type Tabulator = GongTabulator

type GongExcelizeTabulator struct {
	f *excelize.File
}

type ExcelizeTabulator = GongExcelizeTabulator

func (tab *GongExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *GongExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *GongExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *GongExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type GongstructPtr](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type GongstructPtr](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
